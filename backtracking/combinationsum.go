package backtracking

// Given an array of distinct integers candidates and a target integer target,
// return a list of all unique combinations of candidates where the chosen numbers sum to target.
// You may return the combinations in any order.

// The same number may be chosen from candidates an unlimited number of times.
// Two combinations are unique if the frequency of at least one of the chosen numbers is different.

// The test cases are generated such that the number of unique combinations that sum up to target
// is less than 150 combinations for the given input.

// Example 1:
// Input: candidates = [2,3,6,7], target = 7
// Output: [[2,2,3],[7]]
// Explanation:
// 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
// 7 is a candidate, and 7 = 7.
// These are the only two combinations.

func CombinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	backtrackcs(candidates, target, []int{}, 0, 0, &ans)
	return ans
}

func backtrackcs(candidates []int, target int, path []int, start int, curr int, ans *[][]int) {
	if curr == target {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}
	for i := start; i < len(candidates); i++ {
		num := candidates[i]
		if curr+num <= target {
			path = append(path, num)
			backtrackcs(candidates, target, path, i, curr+num, ans)
			path = path[:len(path)-1]
		}
	}
}
