// Package lru provides common LRU cache for go
// implements set and get method
// with hashmap + double link list
package lru

import (
	"sync"
)

type Value interface {
	String() (string, bool)
}

type value struct {
	v interface{}
}

func (v value) String() (string, bool) {
	s, ok := v.v.(string)
	return s, ok
}

type Node struct {
	// key must be string
	key string
	// value can be any type
	value value
	prev  *Node
	next  *Node
}

type LRUCache struct {
	cache map[string]*Node
	sync.Mutex
	// max capacity
	capacity uint64
	// head node
	head *Node
	// tail node
	tail *Node
}

func NewCache(cap uint64) *LRUCache {
	lc := new(LRUCache)
	lc.capacity = cap
	lc.cache = make(map[string]*Node)
	lc.head = new(Node)
	lc.tail = new(Node)
	lc.head.next = lc.tail
	lc.tail.prev = lc.head
	return lc
}

func (lc *LRUCache) addNode(node *Node) {
	node.prev = lc.head
	node.next = lc.head.next

	lc.head.next.prev = node
	lc.head.next = node
}

func (lc *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lc *LRUCache) moveToHead(node *Node) {
	lc.removeNode(node)
	lc.addNode(node)
}

func (lc *LRUCache) popTail() *Node {
	tail := lc.tail.prev
	lc.removeNode(tail)
	return tail
}

func (lc *LRUCache) Get(key string) Value {
	lc.Lock()
	defer lc.Unlock()
	// check exist
	if node, ok := lc.cache[key]; ok {
		lc.moveToHead(node)
		return node.value
	}

	return nil
}

func (lc *LRUCache) Set(key string, v interface{}) error {
	lc.Lock()
	defer lc.Unlock()

	if node, ok := lc.cache[key]; ok {
		node.value = value{v}
		lc.moveToHead(node)
	} else {
		// check cap
		node = &Node{
			key:   key,
			value: value{v},
		}
		lc.cache[key] = node
		lc.addNode(node)

		if len(lc.cache) > int(lc.capacity) {
			// pop tail
			tail := lc.popTail()
			delete(lc.cache, tail.key)
		}
	}

	return nil
}
