package arraysandhashing

func IsSubArray(main []int, sub []int) bool {
	if len(sub) == 0 {
		return true
	}
	if len(sub) > len(main) {
		return false
	}
	for i := 0; i <= len(main)-len(sub); i++ {
		match := true
		for j := range sub {
			if main[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
