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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTokenizer(t *testing.T) {
	tk := newTokenizer("input")
	assert.NotNil(t, tk)
}

func TestParseTokensOK(t *testing.T) {
	cs := []struct {
		Input  string
		Tokens []Token
	}{
		{
			Input: "test==1",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalEqual, Value: "=="},
				{ID: Match, Value: "1"},
				{ID: EOF},
			},
		},
		{
			Input: "test==1||spec=~/one/",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalEqual, Value: "=="},
				{ID: Match, Value: "1"},
				{ID: LogicalOr, Value: "||"},
				{ID: Expr, Value: "spec"},
				{ID: LogicalRegex, Value: "=~"},
				{ID: Match, Value: "/one/"},
				{ID: EOF},
			},
		},
	}
	for i, x := range cs {
		var index = 0
		for item := range newTokenizer(x.Input) {
			checkToken(t, i, index, x.Tokens, item)
			index++
		}
	}
}

func TestParseTokensInvalid(t *testing.T) {
	cs := []struct {
		Input  string
		Tokens []Token
	}{
		{
			Input: "test)=1",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: CloseStatement, Value: ")"},
				{ID: LogicalEqual, Value: "="},
				{ID: Match, Value: "1"},
				{ID: EOF},
			},
		},
		{
			Input: "test<1",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalLessThan, Value: "<"},
				{ID: Match, Value: "1"},
				{ID: EOF},
			},
		},
		{
			Input: "test<=1",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalLessThanOrEqual, Value: "<="},
				{ID: Match, Value: "1"},
				{ID: EOF},
			},
		},
		{
			Input: "test&<=1",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test&"},
				{ID: LogicalLessThanOrEqual, Value: "<="},
				{ID: Match, Value: "1"},
				{ID: EOF},
			},
		},
		{
			Input: "test != 1",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalInvert, Value: "!="},
				{ID: Match, Value: "1"},
				{ID: EOF},
			},
		},
		{
			Input: "(test))>=!=9",
			Tokens: []Token{
				{ID: Entry},
				{ID: OpenStatement, Value: "("},
				{ID: Expr, Value: "test"},
				{ID: CloseStatement, Value: ")"},
				{ID: CloseStatement, Value: ")"},
				{ID: LogicalGreaterThanOrEqual, Value: ">="},
				{ID: Match, Value: ""},
				{ID: LogicalInvert, Value: "!="},
				{ID: Match, Value: "9"},
				{ID: EOF},
			},
		},
		{
			Input: "test===9",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalEqual, Value: "=="},
				{ID: Match, Value: ""},
				{ID: LogicalEqual, Value: "="},
				{ID: Match, Value: "9"},
				{ID: EOF},
			},
		},
	}
	for i, x := range cs {
		var index = 0
		for item := range newTokenizer(x.Input) {
			checkToken(t, i, index, x.Tokens, item)
			index++
		}
	}
}

func TestTokenParserTrimSpace(t *testing.T) {
	cs := []struct {
		Input  string
		Tokens []Token
	}{

		{
			Input: "test == 1",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalEqual, Value: "=="},
				{ID: Match, Value: "1"},
				{ID: EOF},
			},
		},
		{
			Input: "test == 12 || test >= 89",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalEqual, Value: "=="},
				{ID: Match, Value: "12"},
				{ID: LogicalOr, Value: "||"},
				{ID: Expr, Value: "test"},
				{ID: LogicalGreaterThanOrEqual, Value: ">="},
				{ID: Match, Value: "89"},
				{ID: EOF},
			},
		},
		{
			Input: "test =~ /test1",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalRegex, Value: "=~"},
				{ID: Match, Value: "/test1"},
				{ID: EOF},
			},
		},
		{
			Input: "test =~ /testjdksjds/ || test >= 89",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalRegex, Value: "=~"},
				{ID: Match, Value: "/testjdksjds/"},
				{ID: LogicalOr, Value: "||"},
				{ID: Expr, Value: "test"},
				{ID: LogicalGreaterThanOrEqual, Value: ">="},
				{ID: Match, Value: "89"},
				{ID: EOF},
			},
		},
	}
	for i, x := range cs {
		var index = 0
		for item := range newTokenizer(x.Input) {
			checkToken(t, i, index, x.Tokens, item)
			index++
		}
	}
}

func TestParseTokensRegexes(t *testing.T) {
	cs := []struct {
		Input  string
		Tokens []Token
	}{
		{
			Input: "test =~ /test/",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalRegex, Value: "=~"},
				{ID: Match, Value: "/test/"},
				{ID: EOF},
			},
		},
		{
			Input: "test =~ / test /",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalRegex, Value: "=~"},
				{ID: Match, Value: "/ test /"},
				{ID: EOF},
			},
		},
		{
			Input: "test =~ / test\\!() /",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalRegex, Value: "=~"},
				{ID: Match, Value: "/ test\\!() /"},
				{ID: EOF},
			},
		},
		{
			Input: "test =~ /test/ && (test == 1 || test > 18)",
			Tokens: []Token{
				{ID: Entry},
				{ID: Expr, Value: "test"},
				{ID: LogicalRegex, Value: "=~"},
				{ID: Match, Value: "/test/"},
				{ID: LogicalAnd, Value: "&&"},
				{ID: OpenStatement, Value: "("},
				{ID: Expr, Value: "test"},
				{ID: LogicalEqual, Value: "=="},
				{ID: Match, Value: "1"},
				{ID: LogicalOr, Value: "||"},
				{ID: Expr, Value: "test"},
				{ID: LogicalGreaterThan, Value: ">"},
				{ID: Match, Value: "18"},
				{ID: CloseStatement, Value: ")"},
				{ID: EOF},
			},
		},
	}
	for i, x := range cs {
		var index = 0
		for item := range newTokenizer(x.Input) {
			checkToken(t, i, index, x.Tokens, item)
			index++
		}
	}
}

func checkToken(t *testing.T, cs, index int, expected []Token, actual Token) bool {
	if index >= len(expected) {
		t.Errorf("case %d, exceeded expected tokens, token: %s", index, actual)
		return false
	}

	assert.Equal(t, expected[index].ID, actual.ID, "case %d (token: %d), expect: %s but got: %s, token: [%s]", cs, index, expected[index].ID.String(), actual.ID.String(), actual.String())
	assert.Equal(t, expected[index].Value, actual.Value, "case %d (token: %d), expect: %s but got: %s", cs, index, expected[index].Value, actual.Value)

	return true
}
