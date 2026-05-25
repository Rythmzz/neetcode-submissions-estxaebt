/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    curr := root 
	arr := [][]int{}
	dfs(curr,&arr,0)
	return arr
}

func dfs(curr *TreeNode, arr *[][]int, level int) {
	if curr == nil {
		return 
	}

	if len(*arr) == level {
		*arr = append(*arr,[]int{})
	}

	(*arr)[level] = append((*arr)[level], curr.Val)

	dfs(curr.Left, arr,level+1)
	dfs(curr.Right,arr, level+1)
}
