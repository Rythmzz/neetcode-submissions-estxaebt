/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list := []int{}

	for list1 != nil {
		list = append(list,list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		list = append(list,list2.Val)
		list2 = list2.Next
	}

	if len(list) == 0 {
		return nil
	}

	sort.Ints(list)
	
	head := &ListNode{
		Val: list[0],
	}

	curr := head

	for i:= 1;i < len(list) ; i++ {
		curr.Next = &ListNode{Val: list[i]}
		curr = curr.Next
	}

	return head
}
