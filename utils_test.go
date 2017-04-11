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

func TestParseIfFloatOk(t *testing.T) {
	cs := []struct {
		Input    string
		Expected interface{}
	}{
		{Input: "1", Expected: 1.0},
		{Input: "1.3", Expected: 1.3},
		{Input: "4", Expected: 4.0},
		{Input: "45", Expected: 45.0},
	}
	for i, x := range cs {
		found, value := parseIfFloat(x.Input)
		assert.True(t, found, "case %d, expected to find a float", i)
		assert.Equal(t, x.Expected, value, "case %d, expected: %v, got: %v", i, x.Expected, value)
	}
}

func TestParseIfFloatBad(t *testing.T) {
	cs := []struct {
		Input    string
		Expected interface{}
	}{
		{Input: "\"1\"", Expected: "1"},
		{Input: "test", Expected: "test"},
		{Input: "james", Expected: "james"},
		{Input: "4g", Expected: "4g"},
	}
	for i, x := range cs {
		found, value := parseIfFloat(x.Input)
		assert.False(t, found, "case %d, expected not to find a float", i)
		assert.Equal(t, x.Expected, value, "case %d, expected: %v, got: %v", i, x.Expected, value)
	}
}
