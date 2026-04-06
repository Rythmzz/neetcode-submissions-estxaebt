/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	curr := root
	length := findLength(curr)
    list := make([][]int, length)
	appendNode(root,list,0)
	return list

}

func findLength(curr *TreeNode) int {
	if curr == nil {
		return 0
	}

	left:= findLength(curr.Left)
	right:= findLength(curr.Right)

	return 1+max(left,right)
}

func appendNode(curr *TreeNode, result [][]int, offset int) {
	if curr == nil {
		return 
	} 

	result[offset] = append(result[offset],curr.Val)

	appendNode(curr.Left,result,offset+1)
	appendNode(curr.Right,result,offset+1)
}
