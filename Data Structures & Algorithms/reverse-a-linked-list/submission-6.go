/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var l *ListNode
    curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = l
		l = curr
		curr = next
	}

	return l
}
