package linkedlist

func Reverse(head *ListNode) *ListNode {

	prev := (*ListNode)(nil)
	curr := head

	for curr != nil {
		// don't loose next for iterator
		next := curr.next
		// set next equal to previous
		curr.next = prev
		// set previous equal to current
		prev = curr
		// iterator
		curr = next
	}

	return prev
}
