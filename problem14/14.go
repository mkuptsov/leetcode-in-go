package problem14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; ; i++ {
		last := byte(0)
		for _, word := range strs {
			if len(word) <= i {
				return word
			}
			last = strs[0][i]
			if word[i] != last {
				return word[:i]
			}
		}
	}
}
