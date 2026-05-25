/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    curr := root
	invert(curr)
	return root
}

func invert(curr *TreeNode) {
	if curr == nil {
		return
	}

	if curr.Left != nil || curr.Right != nil {
		curr.Left, curr.Right = curr.Right, curr.Left
	}

	invert(curr.Left)
	invert(curr.Right)
}
