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

func TestNewLexer(t *testing.T) {
	assert.NotNil(t, New("test == 1"))
}

func TestAddTokenListener(t *testing.T) {
	l := New("test == 1")
	assert.Empty(t, l.listener)
	l.AddTokenListener(make(TokenChannel, 0))
	assert.NotEmpty(t, l.listener)
	assert.Equal(t, 1, len(l.listener))
}

func TestHasListeners(t *testing.T) {
	l := New("test == 1")
	assert.NotNil(t, l)
	assert.False(t, l.haveListeners())
	l.AddTokenListener(make(TokenChannel, 0))
	assert.True(t, l.haveListeners())
}

func TestParsingRulesBad(t *testing.T) {
	cs := []struct {
		Input string // is the input expression
	}{
		{Input: "test"},
		{Input: "test == 1"},
	}
	for i, c := range cs {
		assert.Error(t, New(c.Input).Parse(), "case %d should have thrown an error", i)
	}
}

func TestIsTokenOk(t *testing.T) {
	cs := []struct {
		Token  Token
		Filter []TokenID
	}{
		{
			Token:  Token{ID: Expr},
			Filter: []TokenID{Expr, OpenStatement},
		},
		{
			Token:  Token{ID: CloseStatement},
			Filter: []TokenID{Expr, OpenStatement, Match, CloseStatement},
		},
	}
	for i, c := range cs {
		assert.True(t, isToken(c.Token, c.Filter), "case %d, should have been true", i)
	}
}

func TestIsTokenBad(t *testing.T) {
	cs := []struct {
		Token  Token
		Filter []TokenID
	}{
		{
			Token:  Token{ID: Expr},
			Filter: []TokenID{OpenStatement, Match, LogicalAnd},
		},
		{
			Token:  Token{ID: CloseStatement},
			Filter: []TokenID{Expr, OpenStatement, Match},
		},
	}
	for i, c := range cs {
		assert.False(t, isToken(c.Token, c.Filter), "case %d, should have been false", i)
	}
}
