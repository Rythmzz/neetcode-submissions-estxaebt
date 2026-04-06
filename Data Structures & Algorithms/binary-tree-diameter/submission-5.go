/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    mD := 0
	findDiameter(root,&mD)

	return mD
}

func findDiameter(curr *TreeNode, mD *int) int {
	if curr == nil {
		return 0
	}

	leftDepth := findDiameter(curr.Left, mD)
	rightDepth := findDiameter(curr.Right,mD)

	if leftDepth + rightDepth > *mD {
		*mD = leftDepth + rightDepth
	}

	return 1 + max(leftDepth,rightDepth)

}
