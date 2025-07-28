package linkedlist

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]
Explanation: 0 + 0 = 0.

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
Explanation: 9999999 + 9999 = 10009998

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// PrintList(l1, "l1")
	// PrintList(l2, "l2")
	dummy := &ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		num1 := 0
		if l1 != nil {
			num1 = l1.val
		}
		num2 := 0
		if l2 != nil {
			num2 = l2.val
		}

		sum := num1 + num2 + carry
		carry = sum / 10
		val := sum % 10

		// println("num1:", num1, "num2:", num2)
		// println("sum:", sum, "carry:", carry, "val:", val)

		curr.next = &ListNode{val: val}
		curr = curr.next
		if l1 != nil {
			l1 = l1.next
		}
		if l2 != nil {
			l2 = l2.next
		}
	}
	return dummy.next
}
