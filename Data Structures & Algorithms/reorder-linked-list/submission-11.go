/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	s,f := head,head

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	var sc *ListNode
	curr := s.Next
	for curr != nil {
		next := curr.Next
		curr.Next = sc
		sc = curr
		curr = next
	}

	fc := head
	s.Next = nil

	for fc != nil && sc != nil {
		fcNext := fc.Next
		scNext := sc.Next

		fc.Next = sc

		if fcNext != nil {
			sc.Next = fcNext
		}

		fc = fcNext
		sc = scNext
	}


}
