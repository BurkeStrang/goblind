package dynamicprogramming

// Given a string s, return the number of palindromic substrings in it.
// A string is a palindrome when it reads the same backward as forward.
// A substring is a contiguous sequence of characters within the string.

// Example 1:
// Input: s = "abc"
// Output: 3
// Explanation: Three palindromic strings: "a", "b", "c".

func CountSubstrings(s string) int {
	count := 0
	for i := range len(s) {
		// odd length palindrome
		count += CountPalindromes(s, i, i)
		// even length palindrome
		count += CountPalindromes(s, i, i+1)
	}
	return count
}

// expand around center
func CountPalindromes(s string, l, h int) int {
	count := 0
	for l >= 0 && h < len(s) && s[l] == s[h] {
		l--
		h++
		count++
	}
	return count
}
