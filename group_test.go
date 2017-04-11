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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAddWhenEmpty(t *testing.T) {
	st := new(Group)
	assert.Nil(t, st.Expression)
	st.Add()
	assert.NotNil(t, st.Expression)
	assert.Nil(t, st.Expression.Next)
}

func TestGetCurrentExpression(t *testing.T) {
	st := new(Group)
	assert.NotNil(t, st)
	assert.Nil(t, st.Last())
	assert.NotNil(t, st.Current())
}

func TestGroupAdd(t *testing.T) {
	st := new(Group)
	st.Add()
	assert.NotNil(t, st.Expression)
	st.Add()
	assert.NotNil(t, st.Expression.Next)
}

func TestGroupLastWhenEmpty(t *testing.T) {
	st := new(Group)
	assert.Nil(t, st.Last())
}

func TestGroupLast(t *testing.T) {
	st := new(Group)
	assert.Nil(t, st.Last())
	e := st.Add()
	e.Match = "test"
	assert.NotNil(t, st.Last())
	assert.Equal(t, "test", st.Last().Match)
}

func TestGroupLastWithMany(t *testing.T) {
	st := new(Group)
	for i := 1; i <= 5; i++ {
		st.Add().Match = fmt.Sprintf("test%d", i)
	}
	assert.Equal(t, "test5", st.Last().Match)
}

func TestGroupSizeWhenEmpty(t *testing.T) {
	assert.Equal(t, 0, new(Group).Size())
}

func TestGroupSize(t *testing.T) {
	st := new(Group)
	for i := 0; i < 5; i++ {
		st.Add().Match = fmt.Sprintf("test%d", i)
	}
	assert.Equal(t, 5, st.Size())
}
