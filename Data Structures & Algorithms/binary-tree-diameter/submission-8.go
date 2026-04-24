/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	dfs(root,&maxDiameter)
	return maxDiameter
}

func dfs(curr *TreeNode, maxDiameter *int) int {
	if curr == nil {
		return 0
	}

	leftDepth := dfs(curr.Left,maxDiameter)
	rightDepth := dfs(curr.Right,maxDiameter)

	if leftDepth + rightDepth > *maxDiameter {
		*maxDiameter = leftDepth + rightDepth
	}

	return 1 + max(leftDepth,rightDepth)
}
