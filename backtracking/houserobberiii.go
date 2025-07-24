package backtracking

func Rob(root *TreeNode) int {
	return getMaxRobAmount(root)
}

func getMaxRobAmount(node *TreeNode) int {
	if node == nil {
		return 0
	}

	rob := node.Val
	if node.Left != nil {
		rob += getMaxRobAmount(node.Left.Left) + getMaxRobAmount(node.Left.Right)
	}
	if node.Right != nil {
		rob += getMaxRobAmount(node.Right.Left) + getMaxRobAmount(node.Right.Right)
	}

	return max(rob, getMaxRobAmount(node.Left)+getMaxRobAmount(node.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
