
package backtracking

import "unicode"


var (
	IsLetter = unicode.IsLetter
	IsUpper  = unicode.IsUpper
	ToUpper  = unicode.ToUpper
	ToLower  = unicode.ToLower
)

func toggleCase(r rune) rune {
	if IsUpper(r) {
		return ToLower(r)
	}
	return ToUpper(r)
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
