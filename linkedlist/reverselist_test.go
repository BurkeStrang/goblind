
package linkedlist

import (
	"reflect"
	"testing"
)


func TestReverse(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		head := BuildList(tt.input)
		reversed := Reverse(head)
		got := listToSlice(reversed)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("Reverse(%v) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}
