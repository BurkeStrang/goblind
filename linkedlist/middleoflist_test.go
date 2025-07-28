package linkedlist

import (
	"testing"
	"reflect"
)


func buildListVals(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.next = &ListNode{val: v}
		curr = curr.next
	}
	return head
}

func listVals(head *ListNode) []int {
	var vals []int
	for head != nil {
		vals = append(vals, head.val)
		head = head.next
	}
	return vals
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6}},
		{[]int{1}, []int{1}},
	}

	for _, tt := range tests {
		head := buildListVals(tt.input)
		got := middleNode(head)
		gotVals := listVals(got)
		if !reflect.DeepEqual(gotVals, tt.expected) {
			t.Errorf("middleNode(%v) = %v, want %v", tt.input, gotVals, tt.expected)
		}
	}
}
