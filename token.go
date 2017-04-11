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

import "fmt"

// String returns the string representaton of the token id type
func (t TokenID) String() string {
	switch t {
	case OpenGroup:
		return "("
	case CloseGroup:
		return ")"
	case LogicalOr:
		return "||"
	case LogicalRegex:
		return "=~"
	case LogicalLessThanOrEqual:
		return "<="
	case LogicalLessThan:
		return "<"
	case LogicalGreaterThanOrEqual:
		return ">="
	case LogicalGreaterThan:
		return ">"
	case LogicalAnd:
		return "&&"
	case LogicalEqual:
		return "=="
	case LogicalInvert:
		return "!="
	case Match:
		return "MATCH"
	case Expr:
		return "EXPR"
	case EOF:
		return "END"
	case Entry:
		return "BEGIN"
	}

	return "Unknown"
}

func (t Token) String() string {
	return fmt.Sprintf("type: '%s', value: '%s'", t.ID.String(), t.Value)
}
