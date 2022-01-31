package reverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	Reverse(&arr)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, arr)

	arr = []int{1}
	Reverse(&arr)
	assert.Equal(t, []int{1}, arr)

	arr = []int{}
	Reverse(&arr)
	assert.Equal(t, []int{}, arr)
}
