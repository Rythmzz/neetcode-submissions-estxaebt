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
	for i:= 0; i< m-1;i++ {
		curr = curr.Next
	}

	curr.Next = curr.Next.Next
	return head
}
