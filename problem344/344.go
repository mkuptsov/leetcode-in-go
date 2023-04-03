package problem344

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < len(s)/2 {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
