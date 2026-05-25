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

	lD := max(0,dfs(curr.Left,maxCurrent))
	rD := max(0,dfs(curr.Right,maxCurrent))

	sum := lD + rD + curr.Val
	*maxCurrent = max(*maxCurrent,sum)

	return curr.Val + max(lD,rD)
}
