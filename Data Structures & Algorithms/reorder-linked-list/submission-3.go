/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
    slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var secondHalf *ListNode
	curr := slow.Next
	
	for curr != nil {
		next := curr.Next
		curr.Next = secondHalf
		secondHalf = curr
		curr =  next
	}

	slow.Next = nil
	firstHalf := head

	for firstHalf != nil && secondHalf != nil {
		fNext := firstHalf.Next
		sNext := secondHalf.Next

		firstHalf.Next = secondHalf

		if fNext != nil {
			secondHalf.Next = fNext
		}

		firstHalf = fNext
		secondHalf = sNext
	}


}
