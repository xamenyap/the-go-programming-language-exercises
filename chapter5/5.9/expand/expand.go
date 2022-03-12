package expand

import (
	"strings"
)

func expand(s string, f func(string) string) string {
	n := f("foo")

	return strings.Replace(s, "$foo", n, -1)
}
