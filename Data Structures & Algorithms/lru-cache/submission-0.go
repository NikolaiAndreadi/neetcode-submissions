type node struct {
	next *node
	prev *node
	key int
	val int
}

type LRUCache struct {
    cache map[int]*node
	capacity int
	head *node
	tail *node
}

func Constructor(capacity int) LRUCache {
    head, tail := &node{}, &node{}
	head.next = tail
	tail.prev = head
	return LRUCache {
		cache: make(map[int]*node),
		capacity: capacity,
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) extract(node *node) {
	node.prev.next, node.next.prev = node.next, node.prev
}

func (this *LRUCache) append(node *node) {
    node.prev, node.next = this.head, this.head.next
    this.head.next.prev, this.head.next = node, node
}

func (this *LRUCache) touch(node *node) {
	this.extract(node)
	this.append(node)
}

func (this *LRUCache) Get(key int) int {
    nod, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.touch(nod)
	return nod.val
}

func (this *LRUCache) Put(key int, value int) {
    nod, ok := this.cache[key]
	if ok {
		nod.val = value
		this.touch(nod)
		return
	}
	nod = &node{key: key, val: value}
	this.append(nod)
	this.cache[key] = nod
	if len(this.cache) > this.capacity {
		goner := this.tail.prev
		this.extract(goner)
		delete(this.cache, goner.key)
	}
}
