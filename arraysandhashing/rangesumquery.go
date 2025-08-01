
package arraysandhashing
// Given an integer array nums, handle multiple queries of the following type:
//
// Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
// Implement the NumArray class:
//
// NumArray(int[] nums) Initializes the object with the integer array nums.
// int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive
// (i.e. nums[left] + nums[left + 1] + ... + nums[right]).
//
// Example 1:
// Input
// ["NumArray", "sumRange", "sumRange", "sumRange"]
// [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
// Output
// [null, 1, -1, -3]
//
// Explanation
// NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
// numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
// numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
// numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3



type NumArray struct {
	sums []int
}


func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	if len(nums) > 0 {
		sums[0] = nums[0]
		for i := 1; i < len(nums); i++ {
			sums[i] = sums[i-1] + nums[i]
		}
	}
	return NumArray{sums: sums}
}

func (na *NumArray) SumRange(left, right int) int {
	if left == 0 {
		return na.sums[right]
	}
	return na.sums[right] - na.sums[left-1]
}
