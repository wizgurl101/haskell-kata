package stringkata

import (
	"reflect"
	"testing"
)

// tests are all passing when some of them should failed
func TestVaildParentheses(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		{
			name:     "should return false",
			args:     args{""},
			expected: false,
		},
		{
			name:     "should return true",
			args:     args{"()"},
			expected: true,
		},
		{
			name:     "should return true",
			args:     args{"()[]{}"},
			expected: true,
		},
		{
			name:     "should return false",
			args:     args{"(]"},
			expected: false,
		},
		{
			name:     "should return true",
			args:     args{"[()]"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VaildParentheses(tt.args.input)
			reflect.DeepEqual(got, tt.expected)
		})
	}
}
