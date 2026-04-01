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
    head := lists[0]

	for i:= 1; i < len(lists);i++{
		curr := head
		mergeTwoLinked(curr,lists[i])
	}

	for curr1 := head; curr1 != nil; curr1 = curr1.Next {
		for curr2 := curr1.Next; curr2 != nil; curr2 = curr2.Next {
			if curr1.Val > curr2.Val {
				temp := curr1.Val
				curr1.Val = curr2.Val
				curr2.Val = temp
			}
		}
	}

	return head

}

func mergeTwoLinked(l1,l2 *ListNode){
	for l1 != nil && l2 != nil {
		l1Next := l1.Next
		l2Next := l2.Next

		l1.Next = l2

		if l1Next != nil {
			l2.Next = l1Next
		}

		l1 = l1Next
		l2 = l2Next
	}
}
