/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}

	leftHigh := getHigh(root.Left)
	rightHigh := getHigh(root.Right)

	if abs(rightHigh - leftHigh) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getHigh(curr *TreeNode) int {
	if curr == nil {
		return 0
	}

	return 1 + max(getHigh(curr.Left),getHigh(curr.Right))
}

func abs(n int) int{
	if n < 0 {
		return -n
	}

	return n
}
