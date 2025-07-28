package linkedlist

import (
	"testing"
)

func TestMergeTwoList(t *testing.T) {
	tests := []struct {
		list1 []int
		list2 []int
		want  []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{0}, []int{0}},
		{[]int{2}, []int{1}, []int{1, 2}},
	}

	for _, tt := range tests {
		l1 := BuildList(tt.list1)
		l2 := BuildList(tt.list2)
		got := MergeTwoList(l1, l2)
		if !listsEqual(got, BuildList(tt.want)) {
			t.Errorf("list not order correctly")
		}
	}
}
