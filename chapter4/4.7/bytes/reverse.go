package bytes

import (
	"unicode/utf8"
)

func Reverse(b []byte) {
	for {
		if len(b) == 0 {
			break
		}

		_, firstByteWidth := utf8.DecodeRune(b)
		_, lastByteWidth := utf8.DecodeLastRune(b)

		if firstByteWidth == 0 || lastByteWidth == 0 {
			break
		}

		rotate(b, firstByteWidth)

		remaining := b[:len(b)-firstByteWidth]
		if len(remaining) == 0 {
			break
		}

		rotate(remaining, len(remaining)-lastByteWidth)

		b = b[lastByteWidth : len(b)-firstByteWidth]
	}
}

func rotate(s []byte, x int) {
	reverse(s[:x])
	reverse(s[x:])
	reverse(s)
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
