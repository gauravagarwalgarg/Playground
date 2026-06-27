package main

import (
	"container/list"
	"fmt"
)

/*
  LC 146 - LRU Cache
  Topic: Advanced Data Structures (HashMap + Doubly Linked List)
  Difficulty: Medium
  Time: O(1) get/put | Space: O(capacity)
*/

type entry struct {
	key, val int
}

type LRUCache struct {
	cap   int
	cache map[int]*list.Element
	order *list.List
}

func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{capacity, make(map[int]*list.Element), list.New()}
}

func (lru *LRUCache) Get(key int) int {
	if el, ok := lru.cache[key]; ok {
		lru.order.MoveToFront(el)
		return el.Value.(entry).val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if el, ok := lru.cache[key]; ok {
		lru.order.MoveToFront(el)
		el.Value = entry{key, value}
		return
	}
	if lru.order.Len() == lru.cap {
		back := lru.order.Back()
		lru.order.Remove(back)
		delete(lru.cache, back.Value.(entry).key)
	}
	el := lru.order.PushFront(entry{key, value})
	lru.cache[key] = el
}

func main() {
	lru := ConstructorLRU(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	if lru.Get(1) == 1 {
		fmt.Println("PASS: Get(1) = 1")
	} else {
		fmt.Println("FAIL: Get(1)")
	}
	lru.Put(3, 3) // evicts key 2
	if lru.Get(2) == -1 {
		fmt.Println("PASS: Get(2) = -1 (evicted)")
	} else {
		fmt.Println("FAIL: Get(2) should be -1")
	}
}
