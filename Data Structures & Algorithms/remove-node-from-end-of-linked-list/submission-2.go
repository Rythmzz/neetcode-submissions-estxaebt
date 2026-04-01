/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil || head.Next == nil {
		return nil
	}

	length := 0
	// 1,2,3,4,5,6,7 n = 3
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}

	m := length - n

	if m < 0 {
		return nil
	}
	curr = head
	c := 0
	for curr != nil {
		if c == m {
			curr = curr.Next
			break
		}
		c++
		curr = curr.Next
	}

	currS := head
	c = 0
	if c == m {
		head = head.Next
	}
	for currS != nil {
		c++
		if c == m {
			currS.Next = curr
			break
		}
		currS = currS.Next
	}

	return head

}
