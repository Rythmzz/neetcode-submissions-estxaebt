/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    h, _ := checkBalance(root)
	return h
}

func checkBalance(curr *TreeNode) (bool,int) {
	if curr == nil {
		return true,0
	}

	left, valLeft := checkBalance(curr.Left)
	if !left {
		return false,0
	}

	right, valRight := checkBalance(curr.Right)
	if !right {
		return false,0
	}

	if abs(valRight - valLeft) > 1 {
		return false,0
	}

	return true, 1 + max(valLeft,valRight)

}

func abs(i int) int{
	if i < 0 { return -i }
	return i
}
