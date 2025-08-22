package arraysandhashing

import "unicode"

func CheckIfPanagram(sentence string) bool {
	var alpha [26]bool

	for _, ch := range sentence {
		if unicode.IsLetter(ch) {
			ch = unicode.ToLower(ch)
			idx := ch - 'a'
			alpha[idx] = true
		}

	}
	for i := range 26 {
		if !alpha[i] {
			return false
		}
	}
	return true
}
