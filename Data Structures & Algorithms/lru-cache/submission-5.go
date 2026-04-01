type LRUCache struct {
	Size int
	Capacity int
	Ch map[int]*ListNode
	Head *ListNode
	Tail *ListNode
}

type ListNode struct {
	Key int
	Val int
	Next *ListNode
	Prev *ListNode
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		Size: 0,
		Capacity: capacity,
		Ch: make(map[int]*ListNode),
		Head: nil,
		Tail: nil,
	}
}

func (this *LRUCache) Get(key int) int {
    if node, exists := this.Ch[key]; exists {
		this.MoveToHead(node)
		return node.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exists := this.Ch[key]; exists {
		node.Val = value
		this.MoveToHead(node)
	} else {

		newNode := &ListNode{
			Key: key,
			Val: value,
			Next: nil,
			Prev: nil,
		}

		this.Ch[key] = newNode

		this.AddToHead(newNode)
		this.Size++

		if this.Size > this.Capacity {
			tail := this.RemoveToTail()
			delete(this.Ch,tail.Key)
			this.Size--
		}
	}


}

func (this *LRUCache) AddToHead(node *ListNode){
	node.Prev = nil
	node.Next = this.Head

	if this.Head != nil {
		this.Head.Prev = node
	}
	this.Head = node

	if this.Tail == nil {
		this.Tail = node
	}
}

func (this *LRUCache) RemoveNode(node *ListNode) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		this.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		this.Tail = node.Prev
	}
}

func (this *LRUCache) MoveToHead(node *ListNode){
	this.RemoveNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) RemoveToTail() *ListNode{
	node:= this.Tail
	this.RemoveNode(node)
	return node
}
