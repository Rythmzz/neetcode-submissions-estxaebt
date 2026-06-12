/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    summary := 0
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil || l2 != nil || summary > 0 {
		sum := summary
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		summary = sum/10
		curr.Next = &ListNode{Val: sum%10}
		curr = curr.Next
	}

	return dummy.Next
}
