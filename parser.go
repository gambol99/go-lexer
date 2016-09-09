/*
Copyright 2016 Rohith Jayawardene <gambol99@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package lex

import (
	"io"
	"strings"
)

// tokenizer
type tokenizer struct {
	input      string       // the actual input
	position   int          // the current position in the string
	start      int          // the start of the cursor
	tokenCh    TokenChannel // the channel to send the tokens
	stopSignal bool         // an exit signal
}

type tokenFn func(*tokenizer) tokenFn

// newTokenizer creates a new tokenizer and starts extract tokens from the input
// sending along the channel
func newTokenizer(input string) TokenChannel {
	t := &tokenizer{
		input:   input,
		tokenCh: make(TokenChannel, 10),
	}

	go func() {
		// step: emit the entrance
		t.emit(Entry)
		// step: look and extract the symbols
		for state := insideExpression; state != nil && !t.stopSignal; {
			state = state(t)
		}
		// step: end the token stream
		t.emit(EOF)
		// step: close these channel
		close(t.tokenCh)
	}()

	return t.tokenCh
}

// stop is responsible for stopping the parser from continuing
func (l *tokenizer) stop() {
	l.stopSignal = true
}

func insideExpression(l *tokenizer) tokenFn {
	c, err := l.next()
	if err == io.EOF {
		l.emit(Expr)
		return nil
	}
	switch c {
	case '(':
		return insideLeftBracket
	case ')':
		return insideRightBracket
	case '&':
		return insideLogicalAnd
	case '|':
		return insideLogicalOr
	case '>':
		return insideGreaterThan
	case '<':
		return insideLessThan
	case '=':
		return insideEquality
	case '!':
		return insideInvertEquality
	}

	return insideExpression
}

// insideMatch indicates we are in matching expression i.e. what were comparing to
func insideMatch(l *tokenizer) tokenFn {
	l.unless([]byte{'(', ')', '&', '|', '>', '<', '=', '!'})
	l.emit(Match)

	return insideExpression
}

func insideLessThan(l *tokenizer) tokenFn {
	l.emitBefore(Expr)

	if c := l.peek(); c != '=' {
		l.emit(LogicalLessThan)
	} else {
		l.ignore()
		l.emit(LogicalLessThanOrEqual)
	}

	return insideMatch
}

func insideInvertEquality(l *tokenizer) tokenFn {
	l.emitBefore(Expr)

	if c := l.peek(); c == '=' {
		l.ignore()
		l.emit(LogicalInvert)

		return insideMatch
	}
	l.backup()

	return insideMatch
}

func insideGreaterThan(l *tokenizer) tokenFn {
	l.emitBefore(Expr)

	if c := l.peek(); c != '=' {
		l.emit(LogicalGreaterThan)
	} else {
		l.ignore()
		l.emit(LogicalGreaterThanOrEqual)
	}

	return insideMatch
}

func insideEquality(l *tokenizer) tokenFn {
	l.emitBefore(Expr)

	switch l.peek() {
	case '=':
		l.ignore()
		l.emit(LogicalEqual)
	case '~':
		l.ignore()
		l.emit(LogicalRegex)
		return outsideRegex
	default:
		l.emit(LogicalEqual)
	}

	return insideMatch
}

func outsideRegex(l *tokenizer) tokenFn {
	c, err := l.next()
	if err == io.EOF {
		return nil
	}
	if c != '/' {
		return outsideRegex
	}

	return insideRegex
}

func insideRegex(l *tokenizer) tokenFn {
	c, err := l.next()
	if err == io.EOF {
		l.emit(Match)
		return nil
	}
	switch c {
	case '/':
		if l.previous() != '\\' {
			l.emit(Match)
			return insideExpression
		}
	}

	return insideRegex
}

func insideLeftBracket(l *tokenizer) tokenFn {
	l.emit(OpenStatement)

	return insideExpression
}

func insideRightBracket(l *tokenizer) tokenFn {
	if l.previous() != ')' {
		l.backup()
		l.emit(Expr)
		l.ignore()
	}
	l.emit(CloseStatement)

	return insideExpression
}

func insideLogicalOr(l *tokenizer) tokenFn {
	if l.peek() != '|' {
		return insideExpression
	}
	// step: emit the logical OR
	l.backup()
	l.emit(Expr)
	l.position += len("||")
	l.emit(LogicalOr)

	return insideExpression
}

func insideLogicalAnd(l *tokenizer) tokenFn {
	if l.peek() != '&' {
		return insideExpression
	}
	// step: emit the logical AN
	l.backup()
	l.emit(Expr)
	l.position += len("&&")
	l.emit(LogicalAnd)

	return insideExpression
}

// emit is responsible for emitting the token upstream
func (l *tokenizer) emit(id TokenID) {
	value := strings.TrimSpace(l.input[l.start:l.position])
	if id == Expr && value == "" {
		return
	}

	l.tokenCh <- Token{
		ID:    id,
		Value: strings.TrimSpace(value),
		Start: l.start,
		End:   l.position,
	}
	l.start = l.position
}

// next is responsible for consuming the next character
func (l *tokenizer) next() (byte, error) {
	// step: have we reached the end of the string?
	if l.position >= len(l.input) {
		return ' ', io.EOF
	}
	ch := l.input[l.position]

	// advance the position
	l.position++

	return ch, nil
}

// backup is responsible for shifting back the cursor
func (l *tokenizer) backup() {
	if l.position > 0 {
		l.position--
	}
}

// ignore pushes the cursor forward by one
func (l *tokenizer) ignore() {
	if l.position <= len(l.input) {
		l.position++
	}
}

// emitBefore emits the token previously and moves on
func (l *tokenizer) emitBefore(id TokenID) {
	l.backup()
	l.emit(id)
	l.ignore()
}

// peek returns the next character and moves back
func (l *tokenizer) peek() byte {
	ch, err := l.next()
	if err == io.EOF {
		return 0
	}
	l.position--

	return ch
}

// unless waits until we hit a character defined or end of file
func (l *tokenizer) unless(filter []byte) {
	for {
		c, err := l.next()
		if err == io.EOF {
			break
		}
		for _, x := range filter {
			if c == x {
				l.backup()
				return
			}
		}
	}
}

// prev allows use to look at the previous char
func (l *tokenizer) previous() byte {
	if l.position == 0 {
		return ' '
	}

	return l.input[l.position-2]
}
