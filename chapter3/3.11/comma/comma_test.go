package comma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComma(t *testing.T) {
	assert.Equal(t, comma("-123456789"), "-123,456,789")
	assert.Equal(t, comma("12345678"), "12,345,678")
	assert.Equal(t, comma("12378.9876"), "12,378.9876")
}
