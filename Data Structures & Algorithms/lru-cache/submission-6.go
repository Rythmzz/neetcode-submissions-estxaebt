type LRUCache struct {
    Size int
	Capacity int
	Ch map[int]*Node
	Head *Node
	Tail *Node
}

type Node struct {
	Val int
	Key int
	Next *Node
	Prev *Node
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		Size: 0,
		Capacity: capacity,
		Ch: make(map[int]*Node),
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
		newNode := &Node{
			Val: value,
			Key: key,
			Next: nil,
			Prev: nil,
		}

		this.Ch[key] = newNode	

		this.AddToHead(newNode)
		this.Size++

		if this.Size > this.Capacity {
			oldNode := this.RemoveToTail()
			delete(this.Ch,oldNode.Key)
			this.Size--
		}
		
	}
}

func (this *LRUCache) RemoveToTail() *Node {
	node := this.Tail
	this.RemoveNode(node)
	return node
}

func (this *LRUCache) MoveToHead(node *Node) {
	this.RemoveNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) AddToHead(node *Node) {
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

func (this *LRUCache) RemoveNode(node *Node){
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

