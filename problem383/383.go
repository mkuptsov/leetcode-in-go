package problem383

func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[rune]int, len(magazine))
	for _, letter := range magazine {
		letters[letter] = letters[letter] + 1
	}
	for _, letter := range ransomNote {
		if item, ok := letters[letter]; !ok || item == 0 {
			return false
		}
		letters[letter] = letters[letter] - 1
	}
	return true
}
