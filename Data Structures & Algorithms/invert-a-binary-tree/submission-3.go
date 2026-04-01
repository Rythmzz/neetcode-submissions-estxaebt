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

	PreOrder(curr)

	return root
	
}

func PreOrder(curr *TreeNode) {
	if curr == nil {
		return
	}

	if curr.Left != nil || curr.Right != nil {
		temp := curr.Left
		curr.Left = curr.Right
		curr.Right = temp
	}




	PreOrder(curr.Left)
	PreOrder(curr.Right)
}


