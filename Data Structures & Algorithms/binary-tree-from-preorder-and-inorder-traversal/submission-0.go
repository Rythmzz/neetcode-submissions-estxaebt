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

	for i:= 0 ;i< len(inorder);i++ {
		if preorder[0] == inorder[i] {
			indexRoot = i
			break
		}
	}

	root := &TreeNode{Val: preorder[0]}

	leftInorder := inorder[:indexRoot]
	rightInorder := inorder[indexRoot+1:]

	leftPreorder := preorder[1:len(leftInorder)+1]
	rightPreorder := preorder[len(leftInorder)+1:]

	root.Left = buildTree(leftPreorder,leftInorder)
	root.Right = buildTree(rightPreorder,rightInorder)

	return root
}


