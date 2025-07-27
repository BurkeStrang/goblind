package linkedlist

import "testing"

func TestHasCycle_NoCycle(t *testing.T) {
	// No cycle
	node1 := &ListNode{val: 1}
	node2 := &ListNode{val: 2}
	node3 := &ListNode{val: 3}
	node1.next = node2
	node2.next = node3
	if HasCycle(node1) {
		t.Error("Expected no cycle, got cycle")
	}
}

func TestHasCycle_SimpleCycle(t *testing.T) {
	// Cycle
	node4 := &ListNode{val: 4}
	node5 := &ListNode{val: 5}
	node4.next = node5
	node5.next = node4 // cycle here
	if !HasCycle(node4) {
		t.Error("Expected cycle, got no cycle")
	}
}

func TestHasCycle_LongCycle(t *testing.T) {
	// another cycle
	node6 := &ListNode{val: 6}
	node7 := &ListNode{val: 7}
	node8 := &ListNode{val: 8}
	node6.next = node7
	node7.next = node8
	node8.next = node6 // cycle here

	if !HasCycle(node6) {
		t.Error("Expected cycle, got no cycle")
	}
}
