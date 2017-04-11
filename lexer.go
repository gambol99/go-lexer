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

	// parsingRules is the ruleset defined the ordering of tokens, i.e. what can the token follow
	parsingRules = map[TokenID][]TokenID{
		Entry:                     {},
		OpenGroup:                 {OpenGroup, LogicalAnd, LogicalOr, Entry},
		CloseGroup:                {CloseGroup, Match},
		Expr:                      {OpenGroup, Entry, LogicalAnd, LogicalOr},
		Match:                     {LogicalEqual, LogicalInvert, LogicalGreaterThan, LogicalGreaterThanOrEqual, LogicalLessThan, LogicalLessThanOrEqual},
		LogicalEqual:              {Expr},
		LogicalLessThan:           {Expr},
		LogicalLessThanOrEqual:    {Expr},
		LogicalGreaterThan:        {Expr},
		LogicalGreaterThanOrEqual: {Expr},
		LogicalInvert:             {Expr},
		LogicalRegex:              {Expr},
		LogicalOr:                 {CloseGroup, Match},
		LogicalAnd:                {CloseGroup, Match},
		EOF:                       {Match, CloseGroup},
	}
)

// New is responsible for creating a new lexer
func New(input string) *Lexer {
	return &Lexer{
		input:    input,
		validFn:  validateExpression,
		listener: make([]TokenChannel, 0),
	}
}

// Parse is responsible for parsing the input stream
func (l *Lexer) Parse() (*Group, error) {
	var lastToken Token          // the previous token we got
	var previous, current *Group // a reference to the current Group
	root := &Group{}

	for i := range newTokenizer(l.input) {
		if l.haveListeners() {
			l.emitTokenListener(i) // handle the token listeners
		}
		// step: if we have a previous token check against the ruleset
		if lastToken.ID != Unknown && !validateTokenRules(lastToken.ID, parsingRules[i.ID]) {
			return nil, fmt.Errorf("'%s' found at position: %d cannot follow '%s'", i.Value, i.Start, lastToken.Value)
		}

		// step: add the token the
		switch i.ID {
		case Entry:
			current = root
		case EOF:
		case OpenGroup:
			previous = current
			current.Next = new(Group)
			current = current.Next
		case CloseGroup:
			if previous == nil {
				return nil, fmt.Errorf("')' closed as position: %d was not opened", i.Start)
			}
			current = previous
			previous = nil
		case LogicalAnd:
			switch lastToken.ID {
			case CloseGroup:
				current.Logic = LogicalTypeAnd
			case Match:
				current.Last().Logic = LogicalTypeAnd
			}
		case LogicalOr:
		case Expr:
			if current.Current().Selector != "" {
				current.Add()
			}
			current.Current().Selector = i.Value
		case Match:
			// step: are we supposed to be a regex?
			switch lastToken.ID {
			case LogicalLessThan:
				fallthrough
			case LogicalLessThanOrEqual:
				fallthrough
			case LogicalGreaterThan:
				fallthrough
			case LogicalGreaterThanOrEqual:
				// step: the match MUST be numeric
				found, v := parseIfFloat(i.Value)
				if !found {
					return nil, fmt.Errorf("value: %s at position: %d must be numeric when using less or greater than", i.Value, i.Start)
				}
				current.Last().Match = v
			case LogicalRegex:
				v, err := regexp.Compile(i.Value)
				if err != nil {
					return nil, fmt.Errorf("regex: '%s' at position: %d is invalid", i.Value, i.Start)
				}
				current.Last().Match = v
			case LogicalEqual:
				// step: convert to float if numeric else leave as a string
				_, v := parseIfFloat(i.Value)
				current.Last().Match = v
			default:
				current.Last().Match = i.Value
			}
		case LogicalRegex:
			fallthrough
		case LogicalEqual, LogicalInvert:
			fallthrough
		case LogicalGreaterThan, LogicalGreaterThanOrEqual:
			fallthrough
		case LogicalLessThan, LogicalLessThanOrEqual:
			current.Last().Operation = getOperation(i.ID)
		default:
			panic("invalid token recieved")
		}
		// step: update the previous token
		lastToken = i
	}

	return root, nil
}

// Evaluate is responsible for evaluating the expression
func (l *Lexer) Evaluate() error {
	// step: we create a

	return nil
}

// AddListener adds a listener to the streams of token produced by the parser
func (l *Lexer) AddListener(ch TokenChannel) *Lexer {
	l.listener = append(l.listener, ch)
	return l
}

// haveListeners checks if we have token listeners
func (l *Lexer) haveListeners() bool {
	return len(l.listener) > 0
}

// emitTokenListener is responsible for forwarding the tokens to the listeners
func (l *Lexer) emitTokenListener(token Token) {
	for _, ch := range l.listener {
		go func(c TokenChannel) {
			c <- token
		}(ch)
	}
}

// validateExpression is the default validation function for an expression
func validateExpression(e string) error {
	if !defaultExpr.MatchString(e) {
		return errors.New("invalid expression")
	}
	return nil
}

// validateTokenRules checks a token complies with the ruleset
func validateTokenRules(id TokenID, filter []TokenID) bool {
	for _, x := range filter {
		if id == x {
			return true
		}
	}

	return false
}
