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

	var sc *ListNode
	curr := slow.Next

	for curr != nil {
		next := curr.Next
		curr.Next = sc
		sc = curr
		curr = next
	}

	slow.Next = nil
	fc := head

	for sc != nil && fc != nil {
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
