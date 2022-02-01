package rotate

func Rotate(s []int, x int) {
	ss := make([]int, x)
	copy(ss, s[:x])
	copy(s[:len(s)-x], s[x:])
	s = s[:len(s)-x]
	s = append(s, ss...)
}
