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

	if same(root,subRoot) {
		return true
	}

	return isSubtree(root.Left,subRoot) || isSubtree(root.Right,subRoot)
}

func same(p *TreeNode, q *TreeNode) bool {
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

	return same(p.Left,q.Left) && same(p.Right,q.Right)
}
