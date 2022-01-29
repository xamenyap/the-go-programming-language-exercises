package anagram

func IsAnagram(s string, ss string) bool {
	if len(s) != len(ss) {
		return false
	}

	if s == "" || ss == "" {
		return false
	}

	if s == ss {
		return false
	}

	m := make(map[int32]bool)
	for _, c := range s {
		m[c] = true
	}

	for _, c := range ss {
		if _, ok := m[c]; !ok {
			return false
		}
	}

	return true
}
