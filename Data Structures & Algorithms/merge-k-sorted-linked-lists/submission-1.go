/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	result := lists[0]
	for i:= 1; i < len(lists);i++{
		result = mergeTwoLinked(result,lists[i])
	}

	return result

}

func mergeTwoLinked(l1,l2 *ListNode) *ListNode{
	temp := &ListNode{}
	curr := temp

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			curr.Next = l2
			l2 = l2.Next
		} else {
			curr.Next = l1
			l1 = l1.Next
		}

		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}

	if l2 != nil {
		curr.Next = l2
	}

	return temp.Next
}
