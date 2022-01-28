package comma

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComma(t *testing.T) {
	assert.Equal(t, comma("12345"), commaUsingBuffer("12345"))
	assert.Equal(t, comma("123456"), commaUsingBuffer("123456"))
	assert.Equal(t, comma("123"), commaUsingBuffer("123"))
	assert.Equal(t, comma("1"), commaUsingBuffer("1"))
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = comma(strconv.Itoa(i))
	}
}

func BenchmarkCommaUsingBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = commaUsingBuffer(strconv.Itoa(i))
	}
}
