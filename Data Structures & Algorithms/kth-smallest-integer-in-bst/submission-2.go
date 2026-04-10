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
	result := 0
	dsf(root,&result,&count,k)
	return result
}

func dsf(curr *TreeNode, result, count *int, k int) {
	if curr == nil {
		return
	}

	dsf(curr.Left,result,count,k)
	*count += 1
	if *count == k {
		*result = curr.Val
	}
	dsf(curr.Right,result,count,k)
}
