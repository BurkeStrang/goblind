package linkedlist

// [1,1,2,1,1,2]

func HasCycle(head *ListNode) bool {
	var fast, slow = head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			return true
		}
	}
	return false
}
