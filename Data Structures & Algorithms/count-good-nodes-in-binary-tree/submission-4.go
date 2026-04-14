/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    count := 0
	dfs(root,&count,root.Val)
	return count
}

func dfs(curr *TreeNode, count *int, maxVal int) {
	if curr == nil {
		return
	}

	if curr.Val >= maxVal {
		*count += 1
		maxVal = curr.Val
	}

	dfs(curr.Left,count,maxVal)
	dfs(curr.Right,count,maxVal)
}
