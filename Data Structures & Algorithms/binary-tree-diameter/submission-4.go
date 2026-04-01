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
	findDiameter(root, &maxDiameter)

	return maxDiameter
}

func findDiameter(curr *TreeNode, maxDiameter *int) int {
	if curr == nil {
		return 0
	}

	leftDepth := findDiameter(curr.Left, maxDiameter)
	rightDepth := findDiameter(curr.Right, maxDiameter)

	if leftDepth + rightDepth > *maxDiameter {
		*maxDiameter = leftDepth + rightDepth
	}

	return 1 + max(leftDepth,rightDepth)
}
