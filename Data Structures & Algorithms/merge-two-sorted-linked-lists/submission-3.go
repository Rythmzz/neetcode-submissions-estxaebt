/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    list := list1
	curr := list

	if curr == nil {
		list = list2
	}

	for curr != nil {
		if curr.Next == nil {
			curr.Next = list2
			break
		}
		curr = curr.Next
	}


	for head := list; head != nil; head = head.Next {
		for head2 := head.Next; head2 != nil; head2 = head2.Next{
			if head.Val > head2.Val {
				temp := head.Val
				head.Val = head2.Val
				head2.Val = temp
			}
		}
	}

	return list

}
