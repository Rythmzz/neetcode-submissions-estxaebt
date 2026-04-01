type LRUCache struct {
	Size int
	Capacity int
	Ch map[int]*Node
	Head *Node
	Tail *Node
}

type Node struct {
	Key int
	Val int
	Next *Node
	Prev *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Size:0,
		Capacity:capacity,
		Ch: make(map[int]*Node),
		Head: nil,
		Tail: nil,
	}
}

func (this *LRUCache) Get(key int) int {
    if node,exist := this.Ch[key]; exist {
		this.MoveToHead(node)
		return node.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node,exist := this.Ch[key]; exist {
		node.Val = value
		this.MoveToHead(node)
		return 
	}

	newNode := &Node{
		Key:key,
		Val:value,
		Next: nil,
		Prev: nil,
	}
	
	this.Ch[key] = newNode
	this.AddToHead(newNode)

	this.Size++

	if this.Size > this.Capacity {
		node := this.RemoveTail()
		delete(this.Ch,node.Key)
		this.Size--
	}
}

func (this *LRUCache) AddToHead(node *Node){
	node.Prev = nil
	node.Next = this.Head

	if this.Head != nil{
		this.Head.Prev = node
	}
	this.Head = node

	if this.Tail == nil {
		this.Tail = node
	}

}

func (this *LRUCache) RemoveNode(node *Node){
	if node.Prev != nil {
		node.Prev.Next =  node.Next
	} else {
		this.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		this.Tail = node.Prev
	}
}

func (this *LRUCache) MoveToHead(node *Node){
	this.RemoveNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) RemoveTail() *Node{
	node := this.Tail
	this.RemoveNode(node)
	return node
}
