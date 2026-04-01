/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    isBalance,_ := checkBalance(root)

	return isBalance
}

func checkBalance(curr *TreeNode) (bool,int) {
	if curr == nil {
		return true,0
	}

	left,valLeft:= checkBalance(curr.Left)
	if !left {
		return false,0
	}

	right,valRight := checkBalance(curr.Right)
	if !right {
		return false,0
	}

	if abs(valLeft - valRight) > 1 {
		return false,0
	}

	return true,1 + max(valLeft,valRight)
}

func abs(a int) int{
	if a < 0 {return -a}
	return a
}
