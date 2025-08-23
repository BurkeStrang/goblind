package arraysandhashing

import "unicode"

func CheckIfPanagram(sentence string) bool {
	alphabet := [26]bool{}
	for _, ch := range sentence {
		if unicode.IsLetter(ch) {
			ch = unicode.ToLower(ch)
			index := ch - 'a'
			alphabet[index] = true
		}
	}
	for j := range 26 {
		if alphabet[j] == false {
			return false
		}
	}
	return true
}
