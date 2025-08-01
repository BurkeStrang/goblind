package arraysandhashing

/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
You must write an algorithm that runs in O(n) time.

Example:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
*/
func LongestConsecutiveSequence(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0
	for _, num := range nums {
		// Only start counting if num-1 is not in the set (start of sequence)
		if _, exists := numSet[num-1]; !exists {
			currentNum := num
			count := 1
			for {
				if _, exists := numSet[currentNum+1]; exists {
					currentNum++
					count++
				} else {
					break
				}
			}
			if count > longest {
				longest = count
			}
		}
	}
	return longest
}
