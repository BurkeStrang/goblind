package linkedlist

import "fmt"

func PrintList(head *ListNode, title string) {
	fmt.Print(title,":")
	visited := make(map[*ListNode]bool)
	for node := head; node != nil; node = node.next {
		fmt.Print(node.val, " ")
		if visited[node] {
			fmt.Print("(cycle detected) ")
			break
		}
		visited[node] = true
	}
	fmt.Println()
}

func BuildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.next = &ListNode{val: n}
		curr = curr.next
	}
	return dummy.next
}

func listToSlice(l *ListNode) []int {
	var res []int
	for l != nil {
		res = append(res, l.val)
		l = l.next
	}
	return res
}
