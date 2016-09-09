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

const (
	// NA is an unknown operation
	NA OperationID = iota
	// EQ means equals to
	EQ
	// NE means not equal
	NE
	// GT means greater than
	GT
	// LT means less than
	LT
	// GTE greater than or equal
	GTE
	// LTE less than or equal
	LTE
	// LIKE is a regex
	LIKE
)

const (
	// Unknown means an unknown token type
	Unknown TokenID = iota
	// Entry is the start of the token stream
	Entry
	// EOF is the end of the token stream
	EOF
	// Expr is an expression
	Expr
	// Match is a value to match the expression
	Match
	// OpenStatement is the start of a statement
	OpenStatement
	// CloseStatement is the end of a statement
	CloseStatement
	// LogicalAnd is a logical AND
	LogicalAnd
	// LogicalOr is a logical OR
	LogicalOr
	// LogicalRegex is a logical regex operation
	LogicalRegex
	// LogicalInvert is a logical invert
	LogicalInvert
	// LogicalEqual is a logical equals operation
	LogicalEqual
	// LogicalLessThan means less than
	LogicalLessThan
	// LogicalLessThanOrEqual means less than or equal
	LogicalLessThanOrEqual
	// LogicalGreaterThan means greter than
	LogicalGreaterThan
	// LogicalGreaterThanOrEqual means greater than or equal
	LogicalGreaterThanOrEqual
)
