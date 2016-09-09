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

// Lexer is the lex parser
type Lexer struct {
	// a list of token channels use to send token ok
	listener []TokenChannel
	// is a validation function for the expression
	validFn exprValidFn
	// the input for the lexer
	input string
}

// TokenID is the token type
type TokenID int

// Token is a token found
type Token struct {
	// ID is the type of token we found
	ID TokenID
	// Value is the string value of the token
	Value string
	// Start is the start of the token
	Start int
	// End is end of the token in the input
	End int
}

// TokenChannel is a channel used to send token upstream
type TokenChannel chan Token

// exprValidFn is a function which validates the expression
type exprValidFn func(string) error
