package bytes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	b := []byte("hello world")
	Reverse(b)
	assert.Equal(t, []byte("dlrow olleh"), b)

	b2 := []byte("\t\rfoo\nbar\f")
	Reverse(b2)
	assert.Equal(t, []byte("\frab\noof\r\t"), b2)
}
