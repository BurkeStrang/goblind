package linkedlist

import "fmt"


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
