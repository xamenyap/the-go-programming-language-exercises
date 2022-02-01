package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimAdjacentSpace(t *testing.T) {
	assert.Equal(t, []byte(" foo bar "), TrimAdjacentSpace([]byte("\n\t\t\rfoo\n\rbar\r\n")))
	assert.Equal(t, []byte(" foo "), TrimAdjacentSpace([]byte("\n\t\t\rfoo\n\r")))
	assert.Equal(t, []byte("foo"), TrimAdjacentSpace([]byte("foo")))
	assert.Equal(t, []byte("\tfoo\nbar "), TrimAdjacentSpace([]byte("\tfoo\nbar\r\t")))
}
