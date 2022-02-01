package rotate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	Rotate(s, 2)
	assert.Equal(t, []int{2, 3, 4, 5, 0, 1}, s)

	s2 := []int{9, 4, 6, 2}
	Rotate(s2, 1)
	assert.Equal(t, []int{4, 6, 2, 9}, s2)
}
