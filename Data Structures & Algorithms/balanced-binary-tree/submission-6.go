/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    h,_:= balance(root)
	return h
}

func balance(curr *TreeNode) (bool,int) {
	if curr == nil {
		return true,0
	}

	left, leftDepth := balance(curr.Left)
	if !left {
		return false,0
	}

	right, rightDepth := balance(curr.Right)
	if !right {
		return false,0
	}

	if abs(leftDepth - rightDepth) > 1 {
		return false,0
	}

	return true, 1+max(leftDepth,rightDepth)
}

func abs(n int) int {
	if n < 0 {return -n}
	return n
}
