package backtracking

import "testing"

func TestRob(t *testing.T) {
	// Example 1: [3,2,3,null,3,null,1]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 1},
		},
	}
	if got := Rob(root1); got != 7 {
		t.Errorf("Rob(root1) = %d; want 7", got)
	}

	// Example 2: [3,4,5,1,3,null,1]
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 1},
		},
	}
	if got := Rob(root2); got != 9 {
		t.Errorf("Rob(root2) = %d; want 9", got)
	}
}
