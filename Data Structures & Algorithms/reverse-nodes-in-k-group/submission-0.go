/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    list := [][]int{}

	curr := head
	i := 0
	arr := []int{}
	for curr != nil {
		i++
		arr = append(arr,curr.Val)
		if i == k {
			list = append(list,arr)
			arr = []int{}
			i = 0
		} 
		curr = curr.Next
		
	}

	if len(arr) != 0 {
		list = append(list,arr)
	}

	fmt.Println(list)

	newList := &ListNode{}
	curr = newList

	for i:= 0 ;i< len(list);i++{
		if len(list[i]) == k {
			for j:= len(list[i])-1;j >= 0 ;j--{
				temp := &ListNode{Val:list[i][j]}
				curr.Next = temp
				curr = curr.Next
			}
		} else {
			for j:= 0;j < len(list[i]) ;j++{
				temp := &ListNode{Val:list[i][j]}
				curr.Next = temp
				curr = curr.Next
			}
		}
	}

	

	return newList.Next

}

func reverseNode(node *ListNode) *ListNode {
	var newNode *ListNode
	curr := node
	for curr != nil {
		next := curr.Next
		curr.Next = newNode
		newNode = curr
		curr = next
	}

	return newNode
}
