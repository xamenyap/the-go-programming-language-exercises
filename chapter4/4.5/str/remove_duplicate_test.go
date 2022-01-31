package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicate(t *testing.T) {
	assert.Equal(t, []string{"foo"}, RemoveDuplicate([]string{"foo", "foo"}))
	assert.Equal(t, []string{"foo", "bar"}, RemoveDuplicate([]string{"foo", "foo", "bar"}))
	assert.Equal(t, []string{"foo", "bar", "foo"}, RemoveDuplicate([]string{"foo", "bar", "foo"}))
	assert.Equal(t, []string{"foo", "bar", "foo"}, RemoveDuplicate([]string{"foo", "bar", "foo", "foo"}))
	assert.Equal(t, []string{"foo", "bar"}, RemoveDuplicate([]string{"foo", "foo", "foo", "bar", "bar"}))
	assert.Equal(t, []string{"foo", "bar", "foo"}, RemoveDuplicate([]string{"foo", "bar", "bar", "foo"}))
	assert.Equal(t, []string{"foo"}, RemoveDuplicate([]string{"foo"}))
}
