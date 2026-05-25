/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	curr := root
    count := 0
	dfs(curr,&count,curr.Val)
	return count
}

func dfs(curr *TreeNode, count *int, maxVal int) {
	if curr == nil {
		return
	}

	if maxVal <= curr.Val {
		*count += 1
		maxVal = curr.Val
	}

	dfs(curr.Left,count,maxVal)
	dfs(curr.Right,count,maxVal)
	
}
