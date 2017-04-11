/*
Copyright 2017 Rohith Jayawardene <gambol99@gmail.com>

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
	// LogicalTypeAnd indicates a AND operation
	LogicalTypeAnd LogicType = 1
	// LogicalTypeOr indicates a OR operation
	LogicalTypeOr LogicType = 0
)

// String returns a string representaton of the logical operation
func (l *LogicType) String() string {
	if *l == LogicalTypeAnd {
		return "&&"
	}

	return "||"
}
