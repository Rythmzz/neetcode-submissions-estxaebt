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
	count := 0
	dfs(root,root.Val,&count)

	return count
}

func dfs(curr *TreeNode, rootVal int, count *int) {
	if curr == nil {
		return
	}

	if curr.Val >= rootVal {
		*count += 1
		rootVal = curr.Val
	}

	dfs(curr.Left,rootVal,count)
	dfs(curr.Right,rootVal,count)
}
