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

	for i,p := range inorder {
		if p == preorder[0] {
			indexRoot = i
			break
		}
	}

	curr := &TreeNode{Val:preorder[0]}

	leftInorder := inorder[:indexRoot]
	rightInorder := inorder[indexRoot+1:]

	leftPreorder := preorder[1:len(leftInorder)+1]
	rightPreorder := preorder[len(leftInorder)+1:]

	curr.Left = buildTree(leftPreorder,leftInorder)
	curr.Right = buildTree(rightPreorder,rightInorder)

	return curr
}
