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
	"errors"
	"fmt"
	"regexp"
)

var (
	defaultExpr = regexp.MustCompile("")

	// parsingRules is the ruleset defined the ordering of tokens, i.e. only the following x come before token y
	parsingRules = map[TokenID][]TokenID{
		Entry:                     {},
		OpenStatement:             {OpenStatement, CloseStatement, LogicalAnd, LogicalOr, Entry},
		CloseStatement:            {CloseStatement, Match},
		Expr:                      {OpenStatement, Entry, LogicalAnd, LogicalOr},
		Match:                     {LogicalEqual, LogicalInvert, LogicalGreaterThan, LogicalGreaterThanOrEqual, LogicalLessThan, LogicalLessThanOrEqual},
		LogicalEqual:              {Expr},
		LogicalLessThan:           {Expr},
		LogicalLessThanOrEqual:    {Expr},
		LogicalGreaterThan:        {Expr},
		LogicalGreaterThanOrEqual: {Expr},
		LogicalInvert:             {Expr},
		LogicalRegex:              {Expr},
		LogicalOr:                 {CloseStatement, Expr},
		LogicalAnd:                {CloseStatement, Expr},
		EOF:                       {Match, CloseStatement},
	}
)

//
// New is responsible for creating a new lexer
//
func New(input string) *Lexer {
	return &Lexer{
		input:    input,
		validFn:  validateExpression,
		listener: make([]TokenChannel, 0),
	}
}

// AddTokenListener adds a listener to the streams of token produced by the parser
func (l *Lexer) AddTokenListener(ch TokenChannel) *Lexer {
	l.listener = append(l.listener, ch)
	return l
}

// Parse is responsible for parsing the input stream
func (l *Lexer) Parse() error {
	var previous Token // the previous token we got

	// step: parse the input stream extracting the tokens and pass through the ruleset
	for i := range newTokenizer(l.input) {
		// step: do we have any listeners to the tokens?
		if l.haveListeners() {
			l.handleTokenListener(i)
		}

		// step: if we have a previous token lets check the current token against the ruleset
		if previous.ID != Unknown && !isToken(i, parsingRules[i.ID]) {
			return fmt.Errorf("'%s' at position: %d cannot follow %s", i.Value, i.Start, previous.ID.String())
		}

		// step: add the token the

		// step: update the previous token
		previous = i
	}

	return nil
}

// Evaluate is responsible for evaluating the expression
func (l *Lexer) Evaluate() error {
	// step: we create a

	return nil
}

// haveListeners checks if we have token listeners
func (l *Lexer) haveListeners() bool {
	return len(l.listener) > 0
}

// handleTokenListener is responsible for forwarding the tokens to the listeners
func (l *Lexer) handleTokenListener(token Token) {
	for _, ch := range l.listener {
		// step: we don't allow the listener to block us
		go func(c TokenChannel) {
			c <- token
		}(ch)
	}
}

// the default validation function for an expression
func validateExpression(e string) error {
	if !defaultExpr.MatchString(e) {
		return errors.New("invalid expression")
	}
	return nil
}

// isToken checks the token is with a select group
func isToken(token Token, filter []TokenID) bool {
	for _, x := range filter {
		if token.ID == x {
			return true
		}
	}

	return false
}
