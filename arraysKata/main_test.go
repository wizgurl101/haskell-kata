package arraysKata

import (
	"reflect"
	"testing"
)

/*
Best practices:
- structure our test a test tables
- properly asert the function behavior
- name test cases as descriptive as possible
*/

func TestGetConcatenation(t *testing.T) {
	type args struct {
		input []int
	}

	// test case definition
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		// test cases
		{
			name:     "should return [1, 2, 1, 1, 2, 1]",
			args:     args{[]int{1, 2, 1}},
			expected: []int{1, 2, 1, 1, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetConcatenation(tt.args.input)
			reflect.DeepEqual(got, tt.expected)
		})
	}
}

func TestFindWordsCharacter(t *testing.T) {
	type args struct {
		first_input  []string
		second_input string
	}

	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name:     "should return [0,1]",
			args:     args{[]string{"leet", "code"}, "e"},
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindsWordsCharacter(tt.args.first_input, tt.args.second_input)
			reflect.DeepEqual(result, tt.expected)
		})
	}
}
