package arraysandhasing

// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving the order of characters.
// No two characters may map to the same character, but a character may map to itself.
//
// Example 1:
// Input: s = "egg", t = "add"
// Output: true
//
// Example 2:
// Input: s = "foo", t = "bar"
// Output: false
//
// Example 3:
// Input: s = "paper", t = "title"
// Output: true
//
// Constraints:
// 1 <= s.length <= 5 * 104
// t.length == s.length
// s and t consist of any valid ascii character.

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[rune]rune)
	tToS := make(map[rune]rune)

	for i, sChar := range s {
		tChar := rune(t[i])

		// Check sToT mapping
		if existingT, exists := sToT[sChar]; exists {
			if existingT != tChar {
				return false
			}
		} else {
			sToT[sChar] = tChar
		}

		// Check tToS mapping
		if existingS, exists := tToS[tChar]; exists {
			if existingS != sChar {
				return false
			}
		} else {
			tToS[tChar] = sChar
		}
	}

	return true
}
