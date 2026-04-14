/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    result := []int{}
	rightAppend(root,&result,0)
	return result
}

func rightAppend(curr *TreeNode, result *[]int, level int) {
	if curr == nil {
		return
	}

	if len(*result) == level {
		*result = append(*result,curr.Val)
	}

	rightAppend(curr.Right,result,level+1)
	rightAppend(curr.Left,result,level+1)
}
