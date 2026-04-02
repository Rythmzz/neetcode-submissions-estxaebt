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

	if sameTree(root,subRoot) {
		return true
	}

	return isSubtree(root.Left,subRoot) || isSubtree(root.Right,subRoot)
}

func sameTree(p,q *TreeNode) bool {
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

	return sameTree(p.Left,q.Left) && sameTree(p.Right,q.Right)
}
