type LRUCache struct {
    capacity int
	size int
	cache map[int]*Node
	head *Node
	tail *Node
}

type Node struct {
	key int
	val int
	prev *Node
	next *Node
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		capacity: capacity,
		size: 0,
		cache: map[int]*Node{},
		head: nil,
		tail: nil,
	}
}

func (this *LRUCache) Get(key int) int {
    if node, exist := this.cache[key]; exist {
		this.MoveToHead(node)
		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exist := this.cache[key]; exist {
		node.val = value
		this.MoveToHead(node)
		return
	}

	newNode := &Node{
		key: key,
		val: value,
		prev: nil,
		next: nil,
	}

	this.cache[key] = newNode
	this.AddToHead(newNode)
	this.size++

	if this.size > this.capacity {
		node := this.RemoveToTail()
		delete(this.cache,node.key)
		this.size--
	}
}

func (this *LRUCache) AddToHead(node *Node) {
	node.prev = nil
	node.next = this.head

	if this.head != nil {
		this.head.prev = node
	}
	this.head = node

	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) RemoveNode(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		this.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		this.tail = node.prev
	}
}

func (this *LRUCache) MoveToHead(node *Node) {
	this.RemoveNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) RemoveToTail() *Node{
	node:= this.tail
	this.RemoveNode(node)
	return node
}
