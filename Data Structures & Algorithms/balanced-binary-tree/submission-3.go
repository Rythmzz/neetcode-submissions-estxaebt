/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    check, _:= checkBalance(root)
	return check
}

func checkBalance(curr *TreeNode) (bool,int) {
	if curr == nil {
		return true, 0
	}

	currLeft, heightLeft := checkBalance(curr.Left)
	if !currLeft {
		return false, 0
	}

	currRight, heightRight := checkBalance(curr.Right)
	if !currRight{
		return false, 0
	}

	if abs(heightLeft - heightRight) > 1 {
		return false, 0
	}

	return true, 1 + max(heightLeft,heightRight)


}

func abs(n int) int{
	if n < 0 {
		return -n
	}

	return n
}
