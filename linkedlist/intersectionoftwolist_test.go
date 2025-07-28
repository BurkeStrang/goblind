package linkedlist

import (
	"testing"
)

func buildListArray(vals ...int) *ListNode {
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

func TestGetIntersectionNode(t *testing.T) {
	// Example 1
	// listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], intersection at node with value 8
	common := buildListArray(8, 4, 5)
	headA := &ListNode{val: 4, next: &ListNode{val: 1, next: common}}
	headB := &ListNode{val: 5, next: &ListNode{val: 6, next: &ListNode{val: 1, next: common}}}
	if got := GetIntersectionNode(headA, headB); got != common {
		t.Errorf("Example 1: expected intersection at %v, got %v", common, got)
	}

	// Example 2
	// listA = [1,9,1,2,4], listB = [3,2,4], intersection at node with value 2
	common2 := buildListArray(2, 4)
	headA2 := &ListNode{val: 1, next: &ListNode{val: 9, next: &ListNode{val: 1, next: common2}}}
	headB2 := &ListNode{val: 3, next: common2}
	if got := GetIntersectionNode(headA2, headB2); got != common2 {
		t.Errorf("Example 2: expected intersection at %v, got %v", common2, got)
	}

	// Example 3
	// listA = [2,6,4], listB = [1,5], no intersection
	headA3 := buildListArray(2, 6, 4)
	headB3 := buildListArray(1, 5)
	if got := GetIntersectionNode(headA3, headB3); got != nil {
		t.Errorf("Example 3: expected no intersection, got %v", got)
	}
}
