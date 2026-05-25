/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    value := 0
	count := 0
	dfs(root,&value,&count,k)
	return value
}

func dfs(curr *TreeNode, value, count *int, k int) {
	if curr == nil {
		return
	}

	dfs(curr.Left,value,count,k)
	*count += 1
	if *count == k {
		*value = curr.Val
	}
	dfs(curr.Right,value,count,k)
}