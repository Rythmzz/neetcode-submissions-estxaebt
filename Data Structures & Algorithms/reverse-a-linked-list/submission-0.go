/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    	stack := []int{}
	for node:= head; node != nil; node = node.Next {
		stack = append(stack,node.Val)
	}

	for node:= head; node != nil; node = node.Next {
		node.Val = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	
	return head
}
