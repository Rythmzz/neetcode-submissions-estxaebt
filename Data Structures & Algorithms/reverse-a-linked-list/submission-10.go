/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	curr := head
	var dummy *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = dummy
		dummy = curr
		curr = next
	}

	return dummy
}
