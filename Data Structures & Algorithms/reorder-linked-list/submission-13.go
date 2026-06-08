/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    slow,fast := head,head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var s2 *ListNode
	curr := slow.Next

	for curr != nil {
		next := curr.Next
		curr.Next = s2
		s2 = curr
		curr = next
	}

	slow.Next = nil
	s1 := head

	for s1 != nil && s2 != nil {
		s1Next := s1.Next
		s2Next := s2.Next

		if s2 != nil {
			s1.Next = s2
		}

		if s1Next != nil {
			s2.Next = s1Next
		}

		s1 = s1Next
		s2 = s2Next
	}

}
