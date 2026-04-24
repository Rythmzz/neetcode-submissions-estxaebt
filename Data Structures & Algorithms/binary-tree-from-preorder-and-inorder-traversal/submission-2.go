/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
		return nil
	}

	indexRoot := 0
	for k, i := range inorder {
		if i == preorder[0] {
			indexRoot = k
			break
		}
	}

	curr := &TreeNode{Val: preorder[0]}

	leftInOrder := inorder[:indexRoot]
	rightInOrder := inorder[indexRoot+1:]

	leftPreOrder := preorder[1:len(leftInOrder)+1]
	rightPreOrder := preorder[len(leftInOrder)+1:]

	curr.Left = buildTree(leftPreOrder,leftInOrder)
	curr.Right = buildTree(rightPreOrder,rightInOrder)

	return curr
	
}
