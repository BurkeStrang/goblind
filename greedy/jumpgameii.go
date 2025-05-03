package greedy

func Jump(nums []int) int {
	left, right, res := 0, 0, 0

	for right < len(nums)-1 {
		maxJump := 0
		for i := left; i <= right; i++ {
			if i+nums[i] > maxJump {
				maxJump = i + nums[i]
			}
		}
		left = right + 1
		right = maxJump
		res++
	}

	return res
}
