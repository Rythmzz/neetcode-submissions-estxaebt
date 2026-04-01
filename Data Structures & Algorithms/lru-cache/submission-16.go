type LRUCache struct {
    size int
	capacity int
	cache map[int]*Node
	head *Node
	tail *Node
}

type Node struct {
	val int
	key int
	next *Node
	prev *Node
}

func Constructor(capacity int) LRUCache {
    return LRUCache {
		size: 0,
		capacity: capacity,
		cache: make(map[int]*Node),
		head: nil,
		tail: nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exist := this.cache[key]; exist {
		this.moveToHead(node)
		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exist := this.cache[key]; exist {
		node.val = value
		this.moveToHead(node)
		return
	}

	newNode := &Node{
		val: value,
		key: key,
		next: nil,
		prev: nil,
	}

	this.addToHead(newNode)
	this.cache[key] = newNode
	this.size++

	if this.size > this.capacity {
		oldNode := this.removeToTail()
		delete(this.cache,oldNode.key)
		this.size--
	}
}

func (this *LRUCache) addToHead(node *Node) {
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

func (this *LRUCache) removeNode(node *Node) {
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

func (this *LRUCache) moveToHead(node *Node){
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeToTail() *Node{
	node := this.tail
	this.removeNode(node)
	return node
}
