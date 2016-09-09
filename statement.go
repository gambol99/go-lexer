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

// evaluate is responsible evaluating the statement
func (s *Statement) evaluate() (bool, error) {
	return false, nil
}

// Add adds an expression to the statement
func (s *Statement) Add() *Expression {
	expr := s.Last()
	if expr == nil {
		s.Expression = new(Expression)
		return s.Expression
	}
	expr.Next = new(Expression)

	return expr.Next
}

// Last gets the last expression in the statement
func (s *Statement) Last() *Expression {
	if s.Expression == nil {
		return nil
	}
	cur := s.Expression
	for cur.Next != nil {
		cur = cur.Next
	}

	return cur
}

// Size gets the number of expressions in the statement
func (s *Statement) Size() int {
	count := 0
	for cur := s.Expression; cur != nil; {
		count++
		cur = cur.Next
	}

	return count
}

// String returns a string representation of the Statement
func (s *Statement) String() string {
	return ""
}