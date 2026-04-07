/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    result := [][]int{}
	dfs(root,&result,0)

	return result
}

func dfs(curr *TreeNode, result *[][]int, level int) {
	if curr == nil {
		return
	}

	if len(*result) == level {
		*result = append(*result,[]int{})
	}

	(*result)[level] = append((*result)[level],curr.Val)

	dfs(curr.Left,result,level+1)
	dfs(curr.Right,result,level+1)

}
