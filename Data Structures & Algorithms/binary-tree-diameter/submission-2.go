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

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		} 

		leftDepth := dfs(node.Left)
		rightDepth := dfs(node.Right)

		if leftDepth + rightDepth > maxDiameter {
			maxDiameter = leftDepth + rightDepth
		}

		return 1 + max(leftDepth,rightDepth)
	}

	dfs(root)
	return maxDiameter
}
