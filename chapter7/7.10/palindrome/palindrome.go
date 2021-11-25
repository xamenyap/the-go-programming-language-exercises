package palindrome

import (
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	length := s.Len()
	i := 0
	j := length - 1

	for i < j {
		equal := !s.Less(i, j) && !s.Less(j, i)
		if !equal {
			return false
		}
		i++
		j--
	}

	return true
}
