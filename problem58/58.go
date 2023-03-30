package problem58

import "strings"

func lengthOfLastWord(s string) int {
    var words []string
    s = strings.Trim(s, " ")
    words = strings.Split(s, " ")
    return len(words[len(words) - 1])
}
