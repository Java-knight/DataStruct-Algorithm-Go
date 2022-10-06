package _146_lru_cache

// LRU 缓存
// https://leetcode.cn/problems/lru-cache/

type LRUCache struct {
	cache   *DoubleLinkedList
	findMap map[int]*Node
	cap     int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cap:     capacity,
		findMap: make(map[int]*Node, 0),
		cache:   &DoubleLinkedList{nil, nil, 0},
	}
	lru.cache.initDoubleLinkedList()
	return lru
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.findMap[key]; !ok {
		return -1
	}
	this.make(key)
	return this.findMap[key].value
}

func (this *LRUCache) make(key int) {
	node := this.findMap[key]
	this.cache.remove(node)
	this.cache.addLast(node)
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.findMap[key]; ok {
		this.deleteCache(key)
		this.addCache(key, value)
		return
	}
	if this.cap == this.cache.size {
		this.removeLast()
	}
	this.addCache(key, value)
}

func (this *LRUCache) deleteCache(key int) {
	node := this.findMap[key]
	this.cache.remove(node)
	delete(this.findMap, key)
}

func (this *LRUCache) addCache(key, val int) {
	node := &Node{key, val, nil, nil}
	this.cache.addLast(node)
	this.findMap[key] = node
}

func (this *LRUCache) removeLast() {
	node := this.cache.removeFirst()
	deleteKey := node.key
	delete(this.findMap, deleteKey)
}

type Node struct {
	key, value int
	pre, next  *Node
}

type DoubleLinkedList struct {
	head, tail *Node
	size       int
}

func (list *DoubleLinkedList) initDoubleLinkedList() {
	list.head = &Node{0, 0, nil, nil}
	list.tail = &Node{0, 0, nil, nil}
	list.head.next = list.tail
	list.tail.pre = list.head
	list.size = 0
}

func (list *DoubleLinkedList) addLast(node *Node) {
	node.pre = list.tail.pre
	node.next = list.tail
	list.tail.pre.next = node
	list.tail.pre = node
	list.size++
}

func (list *DoubleLinkedList) remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
	list.size--
}

func (list *DoubleLinkedList) removeFirst() *Node {
	if list.head.next == list.tail {
		return nil
	}
	node := list.head.next
	list.remove(node)
	return node
}
