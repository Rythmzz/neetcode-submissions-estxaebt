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
	dfs(root,root.Val,&count)
	return count
}

func dfs(curr *TreeNode, maxVal int, count *int) {
	if curr == nil {
		return
	}

	if curr.Val >= maxVal {
		*count +=1
		maxVal = curr.Val
	}

	dfs(curr.Left,maxVal,count)
	dfs(curr.Right,maxVal,count)
}
