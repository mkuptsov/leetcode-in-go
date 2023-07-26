package problem242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	letters := map[rune]int{}

	for _, letter := range s {
		letters[letter]++
	}
	for _, letter := range t {
		_, present := letters[letter]
		if !present {
			return false
		}
		letters[letter]--
		if letters[letter] < 0 {
			return false
		}
	}
	return true
}
