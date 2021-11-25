package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}

func (s stringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s stringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		s        stringSlice
		expected bool
	}{
		{
			name:     "empty slice",
			s:        stringSlice{},
			expected: true,
		},
		{
			name:     "1 element",
			s:        stringSlice{"foo"},
			expected: true,
		},
		{
			name:     "2 different elements",
			s:        stringSlice{"foo", "bar"},
			expected: false,
		},
		{
			name:     "2 same elements",
			s:        stringSlice{"foo", "foo"},
			expected: true,
		},
		{
			name:     "3 same elements",
			s:        stringSlice{"foo", "foo", "foo"},
			expected: true,
		},
		{
			name:     "3 elements",
			s:        stringSlice{"foo", "bar", "baz"},
			expected: false,
		},
		{
			name:     "3 elements, first and last are the same",
			s:        stringSlice{"foo", "bar", "foo"},
			expected: true,
		},
		{
			name:     "4 same elements",
			s:        stringSlice{"foo", "foo", "foo", "foo"},
			expected: true,
		},
		{
			name:     "4 elements, first and last are the same, 2nd and 3rd are different",
			s:        stringSlice{"foo", "bar", "baz", "foo"},
			expected: false,
		},
		{
			name:     "5 elements, correct palindrome",
			s:        stringSlice{"foo", "bar", "baz", "bar", "foo"},
			expected: true,
		},
		{
			name:     "5 elements, incorrect palindrome",
			s:        stringSlice{"foo", "bas", "baz", "bar", "foo"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsPalindrome(tc.s)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
