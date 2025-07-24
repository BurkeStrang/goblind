package backtracking

import "unicode"

// alias for unicode
var IsLetter = unicode.IsLetter

// alias for unicode
var ToUpper = unicode.ToUpper

// alias for unicode
var ToLower = unicode.ToLower

// Find all the combinations of a string given that each character can be uppercase or lowercase
// Example:
// input: s = "Ad2"
// output: ["Ad2","ad2","aD2","AD2"]
func LetterCasePermutation(s string) []string {
	res := []string{}
	backtrack([]rune(s), 0, &res)
	return res
}

func backtrack(chars []rune, index int, res *[]string) {
	if index == len(chars) {
		*res = append(*res, string(chars))
		return
	}

	backtrack(chars, index+1, res)

	if IsLetter(chars[index]) {
		chars[index] = toggleCase(chars[index])
		backtrack(chars, index+1, res)
		chars[index] = toggleCase(chars[index]) // revert case
	}
}

func toggleCase(c rune) rune {
	if unicode.IsUpper(c) {
		return unicode.ToLower(c)
	}
	return unicode.ToUpper(c)
	// if c >= 'a' && c <= 'z' {
	// 	return c - 32
	// }
	// if c >= 'A' && c <= 'Z' {
	// 	return c + 32
	// }
	// return c
}
