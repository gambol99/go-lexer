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

// String returns a string representation of the OperationID
func (o *OperationID) String() string {
	switch *o {
	case EQ:
		return "=="
	case NE:
		return "!="
	case GT:
		return ">"
	case LT:
		return "<"
	case GTE:
		return ">="
	case LTE:
		return "<="
	case LIKE:
		return "=~"
	}

	return "unknown"
}

// getOpertation covers the tokenID to the operation ID
func getOperation(id TokenID) OperationID {
	switch id {
	case LogicalEqual:
		return EQ
	case LogicalInvert:
		return NE
	case LogicalGreaterThan:
		return GT
	case LogicalGreaterThanOrEqual:
		return GTE
	case LogicalLessThan:
		return LT
	case LogicalLessThanOrEqual:
		return LTE
	case LogicalRegex:
		return LIKE
	}

	return NA
}
