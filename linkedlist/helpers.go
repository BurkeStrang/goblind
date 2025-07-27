package linkedlist

import "fmt"

func PrintList(head *ListNode, title string) {
	fmt.Print(title)
	fmt.Println()
	for node := head; node != nil; node = node.next {
		fmt.Print(node.val, " ")
		if node.next == head {
			fmt.Print("(cycle detected) ")
			break
		}
	}
	fmt.Println()
}
