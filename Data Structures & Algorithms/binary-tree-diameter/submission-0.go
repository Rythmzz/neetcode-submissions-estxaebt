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

		depthLeft := dfs(node.Left)
		depthRight := dfs(node.Right)

		if depthLeft + depthRight > maxDiameter {
			maxDiameter = depthLeft + depthRight
		}

		return 1 + max(depthLeft,depthRight)

	}

	dfs(root)
	return maxDiameter

	
}
