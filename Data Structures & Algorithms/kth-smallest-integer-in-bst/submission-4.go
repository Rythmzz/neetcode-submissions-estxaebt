/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    count := 0
	value := 0
	dfs(root,&count,&value,k)
	return value
}

func dfs(curr *TreeNode, count, value *int, k int) {
	if curr == nil {
		return
	}

	dfs(curr.Left,count,value,k)
	*count += 1
	if *count == k {
		*value = curr.Val
	}
	dfs(curr.Right,count,value,k)
}
