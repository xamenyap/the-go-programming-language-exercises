package countingwriter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingWriter(t *testing.T) {
	var b strings.Builder
	cw, written := CountingWriter(&b)
	assert.Equal(t, int64(0), *written)

	_, _ = cw.Write([]byte("foo"))
	assert.Equal(t, int64(3), *written)

	_, _ = cw.Write([]byte("bar"))
	assert.Equal(t, int64(6), *written)

	_, _ = cw.Write([]byte("bazz"))
	assert.Equal(t, int64(10), *written)
}
