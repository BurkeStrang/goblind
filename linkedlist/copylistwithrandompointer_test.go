package linkedlist

import (
	"testing"
)

func buildList(nodes [][]any) *ListNode {
	n := len(nodes)
	list := make([]*ListNode, n)
	for i, node := range nodes {
		list[i] = &ListNode{val: node[0].(int)}
	}
	for i, node := range nodes {
		if i < n-1 {
			list[i].next = list[i+1]
		}
		if node[1] != nil {
			list[i].random = list[node[1].(int)]
		}
	}
	if n > 0 {
		return list[0]
	}
	return nil
}

func listsEqual(a, b *ListNode) bool {
	nodeA, nodeB := a, b
	for nodeA != nil && nodeB != nil {
		if nodeA.val != nodeB.val {
			return false
		}
		if (nodeA.random == nil) != (nodeB.random == nil) {
			return false
		}
		if nodeA.random != nil && nodeB.random != nil {
			if nodeA.random.val != nodeB.random.val {
				return false
			}
		}
		nodeA = nodeA.next
		nodeB = nodeB.next
	}
	return nodeA == nil && nodeB == nil
}

func TestCopyListWithRandomPointer(t *testing.T) {
	input := [][]any{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
	head := buildList(input)
	copied := CopyListWithRandomPointer(head)
	if !listsEqual(head, copied) {
		t.Errorf("copied list does not match original")
	}
}
