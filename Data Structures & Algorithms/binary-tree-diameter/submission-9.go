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
	depth(root,&maxDiameter)
	return maxDiameter
}

func depth(curr *TreeNode, maxDiameter *int) int {
	if curr == nil {
		return 0
	}

	leftDepth := depth(curr.Left,maxDiameter)
	rightDepth := depth(curr.Right,maxDiameter)
	
	if leftDepth + rightDepth > *maxDiameter {
		*maxDiameter = leftDepth + rightDepth
	}

	return 1 + max(leftDepth,rightDepth)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
