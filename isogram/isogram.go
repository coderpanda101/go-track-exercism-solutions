package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	wordSlc := strings.Split(word, "")

	for i := 0; i < len(wordSlc); i++ {
		for j := 0; j < len(wordSlc); j++ {
			if wordSlc[i] == " " || wordSlc[j] == " " || j == i || wordSlc[i] == "-" || wordSlc[j] == "-" {
				continue
			} else if wordSlc[i] == wordSlc[j] {
				return false
			}
		}
	}

	return true
}
