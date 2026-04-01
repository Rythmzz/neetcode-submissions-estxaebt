/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    store := map[*Node]*Node{}

	curr := head
	for curr != nil {
		store[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		if curr.Next != nil {
			store[curr].Next = store[curr.Next]
		} 
		
		if curr.Random != nil {
			store[curr].Random = store[curr.Random]
		}
		
		curr = curr.Next
	}

	return store[head]
}
