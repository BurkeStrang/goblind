package backtracking

import "sort"

/*

Given a collection of candidate numbers (candidates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sum to target.
Each number in candidates may only be used once in the combination.
Note: The solution set must not contain duplicate combinations.

Example 1:
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

Example 2:
Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]

Constraints:
1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30

*/

func CombinationSumII(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	dfs(0, target, []int{}, &result, candidates)
	return result
}

func dfs(start, target int, path []int, result *[][]int, candidates []int) {
	if target == 0 {
		comb := make([]int, len(path))
		copy(comb, path)
		*result = append(*result, comb)
		return
	}
	if target < 0 {
		return
	}
	prev := -1
	for i := start; i < len(candidates); i++ {
		if candidates[i] == prev {
			continue
		}
		path = append(path, candidates[i])
		dfs(i+1, target-candidates[i], path, result, candidates)
		path = path[:len(path)-1]
		prev = candidates[i]
	}
}
