/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    _,h := dfs(root)
	if h {
		return true
	}

	return false
}

func dfs(curr *TreeNode) (int,bool) {
	if curr == nil {
		return 0,true
	}

	leftDepth, left := dfs(curr.Left)
	if !left {
		return 0, false
	}

	rightDepth, right := dfs(curr.Right)
	if !right {
		return 0, false
	}

	if abs(rightDepth - leftDepth) > 1 {
		return 0, false
	}

	return 1+max(rightDepth,leftDepth),true

}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
