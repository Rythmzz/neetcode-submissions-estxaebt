/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    count := 0
	curr := head

	for curr != nil {
		count++
		curr = curr.Next
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for count >= k {
		curr = prev.Next
		next := curr.Next

		for i:= 1;i < k;i++ {
			curr.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
			next = curr.Next
		}

		prev = curr
		count -= k
	}

	return dummy.Next


}
