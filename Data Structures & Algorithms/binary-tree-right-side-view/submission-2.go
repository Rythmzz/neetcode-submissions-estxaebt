/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    if root == nil {
		return nil
	}

	list := []int{root.Val}
	findRight(root,&list,2)

	return list
}

func findRight(curr *TreeNode, list *[]int,level int) {
	if curr == nil || (curr.Left == nil && curr.Right == nil) {
		return
	}

	if curr.Left != nil && curr.Right == nil {
		if len(*list) < level {
		*list = append(*list,curr.Left.Val)
		}
		findRight(curr.Left,list,level+1)
	} else if curr.Left == nil && curr.Right != nil{
		if len(*list) < level {
		*list = append(*list,curr.Right.Val)
		}
		findRight(curr.Right,list,level+1)
	} else {
		if len(*list) < level {
		*list = append(*list,curr.Right.Val)
		}
		findRight(curr.Right,list,level+1)
		findRight(curr.Left,list,level+1)
	}
}
