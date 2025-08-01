package arraysandhashing

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		wantLen  int
		wantNums []int
	}{
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{[]int{0,1,2,2,3,0,4,2}, 2, 5, []int{0,1,3,0,4}},
		{[]int{1,1,1}, 1, 0, []int{}},
		{[]int{4,5}, 3, 2, []int{4,5}},
	}

	for _, tt := range tests {
		numsCopy := append([]int(nil), tt.nums...)
		gotLen := RemoveElement(numsCopy, tt.val)
		if gotLen != tt.wantLen || !reflect.DeepEqual(numsCopy[:gotLen], tt.wantNums) {
			t.Errorf("RemoveElement(%v, %d) = %d, %v; want %d, %v",
				tt.nums, tt.val, gotLen, numsCopy[:gotLen], tt.wantLen, tt.wantNums)
		}
	}
}
