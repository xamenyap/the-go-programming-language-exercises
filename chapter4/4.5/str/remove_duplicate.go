package str

func RemoveDuplicate(strs []string) []string {
	if len(strs) < 2 {
		return strs
	}

	i := 0
	j := 1
	for {
		if j == len(strs) {
			break
		}

		if strs[i] == strs[j] {
			copy(strs[j:], strs[j+1:])
			strs = strs[:len(strs)-1]
			continue
		}

		i++
		j++
	}

	return strs
}
