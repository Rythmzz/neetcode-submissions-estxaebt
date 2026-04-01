/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if checkSame(root,subRoot) {
		return true
	}

	return isSubtree(root.Left,subRoot) || isSubtree(root.Right,subRoot)
	
}

func checkSame(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
	} else {
		return false
	}

	return checkSame(p.Left,q.Left) && checkSame(p.Right,q.Right)
}

