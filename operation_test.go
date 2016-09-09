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

func TestOperationString(t *testing.T) {
	cs := []struct {
		ID       OperationID
		Expected string
	}{
		{ID: EQ, Expected: "=="},
		{ID: NE, Expected: "!="},
		{ID: GT, Expected: ">"},
		{ID: LT, Expected: "<"},
		{ID: GTE, Expected: ">="},
		{ID: LTE, Expected: "<="},
		{ID: LIKE, Expected: "=~"},
		{ID: NA, Expected: "unknown"},
	}
	for _, c := range cs {
		assert.Equal(t, c.Expected, c.ID.String())
	}
}

func TestGetOperation(t *testing.T) {
	cs := []struct {
		ID       TokenID
		Expected OperationID
	}{
		{ID: LogicalEqual, Expected: EQ},
		{ID: LogicalInvert, Expected: NE},
		{ID: LogicalGreaterThan, Expected: GT},
		{ID: LogicalGreaterThanOrEqual, Expected: GTE},
		{ID: LogicalLessThan, Expected: LT},
		{ID: LogicalLessThanOrEqual, Expected: LTE},
		{ID: LogicalRegex, Expected: LIKE},
		{ID: Unknown, Expected: NA},
	}
	for _, c := range cs {
		assert.Equal(t, c.Expected, getOperation(c.ID))
	}
}
