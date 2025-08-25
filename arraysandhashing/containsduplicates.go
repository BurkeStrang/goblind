package arraysandhashing

func ContainsDuplicate(nums []int) bool {
	unique := make(map[int]struct{})
	for _, num := range nums {
		if _, exists := unique[num]; exists {
			return true
		} else {
			unique[num] = struct{}{}
		}
	}
	return false
}
