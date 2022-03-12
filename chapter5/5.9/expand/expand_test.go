package expand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		f        func(string) string
		expected string
	}{
		{
			name:  "simplest case",
			input: "$foo",
			f: func(s string) string {
				return "bar"
			},
			expected: "bar",
		},
		{
			name:  "no match",
			input: "$bar",
			f: func(s string) string {
				return "bar"
			},
			expected: "$bar",
		},
		{
			name:  "two matches",
			input: "$foo$bar$foo",
			f: func(s string) string {
				return "baz"
			},
			expected: "baz$barbaz",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := expand(tc.input, tc.f)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
