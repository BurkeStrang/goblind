package linkedlist

// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists into one sorted list.
// The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.

func MergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	curr := res

	for list1 != nil && list2 != nil {
		if list1.val < list2.val {
			curr.next = list1
			list1 = list1.next
		} else {
			curr.next = list2
			list2 = list2.next
		}
		curr = curr.next
	}

	if list1 != nil {
		curr.next = list1
	}
	if list2 != nil {
		curr.next = list2
	}

	return res.next
}
