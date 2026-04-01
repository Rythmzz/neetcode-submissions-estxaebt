/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    seen := map[int]bool{}

	curr := head
	for curr != nil {
		if seen[curr.Val] && curr.Next != nil {
			return true
		}
		seen[curr.Val] = true
		
		curr = curr.Next
	}

	return false
}
