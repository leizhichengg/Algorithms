package main

import (
	"fmt"
)

type Node struct {
	key, value int
	pre        *Node
	next       *Node
}

func NewNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

type LRUCache struct {
	CacheMap   map[int]*Node
	Capacity   int
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	head := NewNode(0, 0)
	tail := NewNode(0, 0)
	head.next = tail
	tail.pre = head
	head.pre = nil
	tail.next = nil
	return LRUCache{
		CacheMap: make(map[int]*Node),
		Capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) deleteNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) pushToHead(node *Node) {
	node.next = this.head.next
	node.next.pre = node
	node.pre = this.head
	this.head.next = node
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.CacheMap[key]; ok {
		value := node.value
		this.deleteNode(node)
		this.pushToHead(node)
		return value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		node := this.CacheMap[key]
		node.value = value
	} else {
		node := NewNode(key, value)
		if len(this.CacheMap) >= this.Capacity {
			delete(this.CacheMap, this.tail.pre.key)
			this.deleteNode(this.tail.pre)
			this.pushToHead(node)
		} else {
			this.pushToHead(node)
		}
		this.CacheMap[key] = node
	}
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
