package comma

import (
	"bytes"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func anotherComma(s string) string {
	commaPos := make(map[int]struct{})
	cur := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur++
		if cur%3 == 0 {
			commaPos[i] = struct{}{}
		}
	}

	var b bytes.Buffer
	for i := 0; i < len(s); i++ {
		if _, ok := commaPos[i]; ok && i != 0 {
			b.WriteByte(',')
		}
		b.WriteByte(s[i])
	}

	return b.String()
}
