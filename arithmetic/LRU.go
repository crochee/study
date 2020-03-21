package arithmetic

import (
	"sync"
)

/*
* @
* @Author:
* @Date: 2020/3/21 17:08
 */
// LRUCache Least Recently Used 最少使用池
type LRUCache struct {
	mux      *sync.Mutex
	list     *DoubleLinkList
	capacity int
}

// 改进数据同步,主要是不同协程改变数据后，数据隔离的问题
// NewLRUCache 生成最少使用池对象
func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		mux:      new(sync.Mutex),
		list:     NewDoubleLinkList(),
		capacity: cap,
	}
}

// Set 添加元素
func (lru *LRUCache) Set(key, value interface{}) {
	if lru.list.size < lru.capacity {
		node := lru.list.Exist(key)
		lru.mux.Lock()
		if node != nil {
			lru.list.Remove(node)
		}
		lru.list.Append(key, value)
		lru.mux.Unlock()
	} else {
		node := lru.list.Exist(key)
		if node == nil {
			node = lru.list.head.next
		}
		lru.mux.Lock()
		lru.list.Remove(node)
		lru.list.Append(key, value)
		lru.mux.Unlock()
	}
}

// Get 获取元素
func (lru *LRUCache) Get(key interface{}) interface{} {
	node := lru.list.Exist(key)
	lru.mux.Lock()
	lru.list.Remove(node)
	lru.list.Append(key, node.value)
	lru.mux.Unlock()
	return node.value
}

// remove 删除元素
func (lru *LRUCache) remove(key interface{}) {
	node := lru.list.Exist(key)
	if node != nil {
		lru.mux.Lock()
		lru.list.Remove(node)
		lru.mux.Unlock()
	}
}

func (lru *LRUCache) All() map[interface{}]interface{} {
	temp := make(map[interface{}]interface{}, lru.capacity)
	node := lru.list.tail
	for node != lru.list.head {
		temp[node.key] = node.value
		node = node.pre
	}
	return temp
}

// DoubleLinkNode 双向链表节点
type DoubleLinkNode struct {
	key   interface{}
	value interface{}
	next  *DoubleLinkNode
	pre   *DoubleLinkNode
}

// DoubleLinkList 双向链表
type DoubleLinkList struct {
	head *DoubleLinkNode
	tail *DoubleLinkNode
	size int
}

// NewDoubleLinkList 生成双向链表对象
func NewDoubleLinkList() *DoubleLinkList {
	var head = &DoubleLinkNode{}
	return &DoubleLinkList{
		head: head,
		tail: head,
		size: 0,
	}
}

// 尾部插入数据
func (dll *DoubleLinkList) Append(key, value interface{}) {
	tail := &DoubleLinkNode{
		key:   key,
		value: value,
		next:  nil,
		pre:   dll.tail,
	}
	dll.tail.next = tail
	dll.tail = tail
	dll.size++
}

// 删除元素
func (dll *DoubleLinkList) Remove(node *DoubleLinkNode) {
	if node.next == nil { // 删除最后一个节点
		dll.tail.key = node.pre.key
		dll.tail.value = node.pre.value
		dll.tail.pre = node.pre
		dll.tail.next = nil
		node = nil
	} else { // 删除中间节点
		node.key = node.next.key
		node.value = node.next.value
		node.next = node.next.next
		node.next.pre = node.pre
	}
	dll.size--
}

func (dll *DoubleLinkList) Exist(key interface{}) *DoubleLinkNode {
	node := dll.tail
	for node != dll.head {
		if node.key == key {
			return node
		}
		node = node.pre
	}
	return nil
}
