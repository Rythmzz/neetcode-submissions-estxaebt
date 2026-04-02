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

	var sc *ListNode
	curr := slow.Next
	for curr != nil {
		next := curr.Next
		curr.Next = sc
		sc = curr
		curr = next
	}

	fc := head
	slow.Next = nil

	for sc != nil && fc != nil {
		fcN := fc.Next
		scN := sc.Next

		fc.Next = sc

		if fcN != nil {
			sc.Next = fcN
		}

		fc = fcN
		sc = scN
	}

}
