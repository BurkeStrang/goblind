package backtracking

// Find all the combinations of a string given that each character can be uppercase or lowercase
// Example:
// input: s = "Ad2"
// output: ["Ad2","ad2","aD2","AD2"]

// LetterCasePermutation returns all combinations of the input string.
// with each letter toggled between lowercase and uppercase.
func LetterCasePermutation(s string) []string {
	var results []string
	runes := []rune(s)
	backtrack(runes, 0, &results)
	return results
}

// backtrack recursively generates case permutations starting from index.
func backtrack(runes []rune, index int, results *[]string) {
	if index == len(runes) {
		*results = append(*results, string(runes))
		return
	}

	// Continue without changing current character
	backtrack(runes, index+1, results)

	// If it's a letter, toggle its case and continue
	if IsLetter(runes[index]) {
		runes[index] = toggleCase(runes[index])
		backtrack(runes, index+1, results)
		// Revert back to original case after recursion
		runes[index] = toggleCase(runes[index])
	}
}
