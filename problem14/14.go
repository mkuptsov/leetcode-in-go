package problem14

import "strings"

func longestCommonPrefix(strs []string) string {
	var prefix strings.Builder
	if len(strs) == 0 {
		return ""
	}
	for i := 0; ; i++ {
		last := byte(0)
		for _, word := range strs {
			if len(word) <= i {
				return prefix.String()
			}
			last = strs[0][i]
			if word[i] != last {
				return prefix.String()[:i]
			}
		}
		prefix.WriteByte(last)
	}
}
