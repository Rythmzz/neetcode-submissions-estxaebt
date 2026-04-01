/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := l1
	list2 := l2
    list := &ListNode{}
	curr := list
	carry := 0

	for list1 != nil || list2 != nil || carry > 0 {
		sum := carry

		if list1 != nil {
			sum += list1.Val
			list1 = list1.Next
		}

		if list2 != nil {
			sum += list2.Val
			list2 = list2.Next
		}

		carry = sum / 10
		curr.Next = &ListNode{Val: sum %10}
		curr = curr.Next
	}

	return list.Next
}
