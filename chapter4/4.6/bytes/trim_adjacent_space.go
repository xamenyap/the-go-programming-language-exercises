package bytes

import (
	"unicode"
)

func TrimAdjacentSpace(b []byte) []byte {
	i := 0
	j := 1
	for {
		if j == len(b) {
			break
		}

		if unicode.IsSpace(rune(b[i])) && unicode.IsSpace(rune(b[j])) {
			b[i] = ' '
			copy(b[j:], b[j+1:])
			b = b[:len(b)-1]

			continue
		}

		i++
		j++
	}

	return b
}
