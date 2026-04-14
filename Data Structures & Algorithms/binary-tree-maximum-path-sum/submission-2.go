/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    maxCurrent := root.Val
	dfs(root,&maxCurrent)
	return maxCurrent
}

func dfs(curr *TreeNode, maxCurrent *int) int {
	if curr == nil {
		return 0
	}

	leftSum := max(0,dfs(curr.Left,maxCurrent))
	rightSum := max(0,dfs(curr.Right,maxCurrent))

	sum := curr.Val + leftSum + rightSum

	*maxCurrent = max(*maxCurrent,sum)

	return curr.Val + max(leftSum,rightSum)
}
