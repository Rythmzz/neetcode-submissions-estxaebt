/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    _,h := balance(root)
	if h {
		return true
	}

	return false
}

func balance(curr *TreeNode) (int,bool) {
	if curr == nil {
		return 0, true
	}

	leftDepth, left := balance(curr.Left)
	if !left {
		return 0, false
	}

	rightDepth, right := balance(curr.Right)
	if !right {
		return 0, false
	}

	if abs(leftDepth - rightDepth) > 1 {
		return 0, false
	}

	return 1 + max(leftDepth,rightDepth), true

}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
