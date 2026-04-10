/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    list := []int{}

	dfs(root,&list)
	fmt.Println(list)


	return list[k-1]

}

func dfs(curr *TreeNode, list *[]int){
	if curr == nil {
		return
	}


	dfs(curr.Left,list)
	*list = append(*list,curr.Val)
	dfs(curr.Right,list)
}
