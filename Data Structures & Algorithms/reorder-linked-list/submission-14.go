/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	curr := slow.Next
	var l2 *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = l2
		l2 = curr
		curr = next
	}

	slow.Next = nil
	l1 := head

	for l1 != nil && l2 != nil {
		l1N := l1.Next
		l2N := l2.Next

		if l2 != nil {
			l1.Next = l2
		}

		if l1N != nil {
			l2.Next = l1N
		}

		l1 = l1N
		l2 = l2N
	}

}
