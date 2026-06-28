package main

import "fmt"

// LRU Cache with doubly-linked list + hash map.
// O(1) Get and Put operations.
// When capacity is exceeded, the least recently used item is evicted.

type node struct {
	key, val   int
	prev, next *node
}

type LRUCache struct {
	capacity int
	cache    map[int]*node
	head     *node // Most recently used (dummy)
	tail     *node // Least recently used (dummy)
}

func NewLRUCache(capacity int) *LRUCache {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*node),
		head:     head,
		tail:     tail,
	}
}

// Get returns the value and marks it as recently used. Returns -1 if not found.
func (l *LRUCache) Get(key int) int {
	if n, ok := l.cache[key]; ok {
		l.remove(n)
		l.insertFront(n)
		return n.val
	}
	return -1
}

// Put inserts or updates a key-value pair. Evicts LRU item if at capacity.
func (l *LRUCache) Put(key, value int) {
	if n, ok := l.cache[key]; ok {
		n.val = value
		l.remove(n)
		l.insertFront(n)
		return
	}

	if len(l.cache) >= l.capacity {
		// Evict least recently used (node before tail)
		lru := l.tail.prev
		l.remove(lru)
		delete(l.cache, lru.key)
	}

	newNode := &node{key: key, val: value}
	l.cache[key] = newNode
	l.insertFront(newNode)
}

// remove detaches a node from the linked list.
func (l *LRUCache) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

// insertFront places a node right after the head (most recently used position).
func (l *LRUCache) insertFront(n *node) {
	n.next = l.head.next
	n.prev = l.head
	l.head.next.prev = n
	l.head.next = n
}

func main() {
	cache := NewLRUCache(3)

	cache.Put(1, 10)
	cache.Put(2, 20)
	cache.Put(3, 30)
	fmt.Printf("Get(1): %d\n", cache.Get(1)) // 10 (moves 1 to front)

	cache.Put(4, 40) // Evicts key 2 (LRU)
	fmt.Printf("Get(2): %d\n", cache.Get(2)) // -1 (evicted)
	fmt.Printf("Get(3): %d\n", cache.Get(3)) // 30
	fmt.Printf("Get(4): %d\n", cache.Get(4)) // 40

	cache.Put(5, 50) // Evicts key 1
	fmt.Printf("Get(1): %d\n", cache.Get(1)) // -1 (evicted)
	fmt.Printf("Get(5): %d\n", cache.Get(5)) // 50
}
