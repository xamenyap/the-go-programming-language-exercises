package anagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	assert.Equal(t, true, IsAnagram("dormitory", "dirtyroom"))
	assert.Equal(t, true, IsAnagram("adecimalpoint", "imadotinplace"))
	assert.Equal(t, true, IsAnagram("snoozealarms", "alasnomorezs"))
	assert.Equal(t, false, IsAnagram("", ""))
	assert.Equal(t, false, IsAnagram("abc", "abc"))
	assert.Equal(t, false, IsAnagram("abc", "ab"))
	assert.Equal(t, false, IsAnagram("tomato", "tamote"))
}
