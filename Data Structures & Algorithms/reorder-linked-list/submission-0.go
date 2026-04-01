/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil && head.Next == nil {
		return
	}

    slow, fast := head,head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var secondHalf *ListNode
	curr := slow.Next
	for curr != nil {
		next := curr.Next
		curr.Next = secondHalf
		secondHalf = curr
		curr = next
	}

	slow.Next = nil
	firstHalf := head

	mergeLinkedList(firstHalf,secondHalf)
}

func mergeLinkedList(l1,l2 *ListNode) {
	for l1 != nil && l2 != nil {
		l1Next := l1.Next
		l2Next := l2.Next

		l1.Next = l2

		if l1Next != nil {
			l2.Next = l1Next
		}

		l1 = l1Next
		l2 = l2Next
	}
}
