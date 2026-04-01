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

	var s *ListNode
	curr := slow.Next
	for curr != nil {
		next := curr.Next
		curr.Next = s
		s = curr
		curr = next
	}

	slow.Next = nil
	f := head

	for s != nil && f != nil {
		fNext := f.Next
		sNext := s.Next

		f.Next = s

		if fNext != nil {
			s.Next = fNext
		}

		f = fNext
		s = sNext
	}

}
