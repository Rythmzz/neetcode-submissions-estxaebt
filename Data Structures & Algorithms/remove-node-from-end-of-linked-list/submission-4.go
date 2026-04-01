/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := 0
	curr := head

	for curr != nil {
		length++
		curr = curr.Next
	}

	m := length - n

	if m == 0 {
		return head.Next
	}

	curr = head
	i := 0
	for curr != nil {
		i++
		if i == m {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}

	return head
}
