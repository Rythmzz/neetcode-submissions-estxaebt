/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    if root == nil {
		return false
	}

	return dfs(root,math.MinInt64,math.MaxInt64)
}

func dfs(curr *TreeNode, minVal,maxVal int) bool {
	if curr == nil {
		return true
	}

	if curr.Val <= minVal || curr.Val >= maxVal {
		return false
	}

	return dfs(curr.Left,minVal,curr.Val) && dfs(curr.Right,curr.Val, maxVal)
	
}
