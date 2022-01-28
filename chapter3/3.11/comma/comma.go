package comma

import (
	"strings"
)

func comma(s string) string {
	var sign string
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		sign = string(s[0])
		s = s[1:]
	}

	var afterIdx string
	if idx := strings.LastIndex(s, "."); idx != -1 {
		afterIdx = s[idx:]
		s = s[:idx]
	}

	return sign + addComma(s) + afterIdx
}

func addComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
