package intset

import (
	"bytes"
	"fmt"
)

const effectiveSize = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/effectiveSize, uint(x%effectiveSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/effectiveSize, uint(x%effectiveSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	maxLen := len(s.words)
	if maxLen < len(t.words) {
		maxLen = len(t.words)
	}

	for i := 0; i < maxLen; i++ {
		if i >= len(s.words) {
			continue
		}

		if i >= len(t.words) {
			s.words[i] = 0
			continue
		}

		s.words[i] &= t.words[i]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] = s.words[i] & (^t.words[i])
		}
	}
}

func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] ^= t.words[i]
		}
	}

	for i := range t.words {
		if i >= len(s.words) {
			s.words = append(s.words, t.words[i])
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < effectiveSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", effectiveSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	sLen := 0
	for _, word := range s.words {
		for i := 0; i < effectiveSize; i++ {
			if word&(1<<uint(i)) != 0 {
				sLen++
			}
		}
	}

	return sLen
}

func (s *IntSet) Remove(x int) {
	word, bit := x/effectiveSize, uint(x%effectiveSize)
	if word > len(s.words) {
		return
	}

	s.words[word] = s.words[word] &^ (1 << bit)
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	cp := new(IntSet)
	cp.words = make([]uint, len(s.words), cap(s.words))
	copy(cp.words, s.words)

	return cp
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *IntSet) Elems() []int {
	elems := make([]int, 0)
	for idx, word := range s.words {
		for i := 0; i < effectiveSize; i++ {
			if word&(1<<uint(i)) != 0 {
				elems = append(elems, idx*effectiveSize+i)
			}
		}
	}

	return elems
}
