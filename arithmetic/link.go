package arithmetic

import "fmt"

/**
* @ Description:
* @Author:
* @Date: 2020/3/10 17:48
 */
type Node struct {
	data interface{}
	next *Node
}

type LinkList struct {
	head *Node
	tail *Node
	size int
}

func NewLink() *LinkList {
	return &LinkList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l *LinkList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkList) Len() int {
	return l.size
}

func (l *LinkList) exit(node *Node) bool {
	p := l.head
	for p != nil {
		if p == node {
			return true
		} else {
			p = p.next
		}
	}
	return false
}
func (l *LinkList) Get(e interface{}) *Node {
	p := l.head
	for p != nil {
		if e == p.data {
			return p
		} else {
			p = p.next
		}
	}
	return nil
}

func (l *LinkList) Append(e interface{}) {
	newNode := &Node{
		data: e,
		next: nil,
	}
	if l.size == 0 { // 空链表时头尾都是新节点
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode // 当前尾节点的下一个节点为新节点，指针指向下一个
		l.tail = newNode      // 更新后的链表的尾节点为新节点，值更新
	}
	l.size++
}

func (l *LinkList) InsertHead(e interface{}) {
	newNode := &Node{
		data: e,
		next: l.head,
	}
	l.head = newNode
	if l.size == 0 {
		l.tail = newNode
	}
	l.size++
}
func (l *LinkList) InsertAfterNode(pre *Node, e interface{}) {
	if l.exit(pre) {
		newNode := &Node{
			data: e,
		}
		if pre.next == nil {
			l.Append(e)
		} else {
			newNode.next = pre.next
			pre.next = newNode
		}
		l.size++
	} else {
		fmt.Println("is not exit ", pre)
	}
}

func (l *LinkList) InsertAfterData(pre, e interface{}) bool {
	p := l.head
	for p != nil {
		if p.data == pre {
			l.InsertAfterNode(p, e)
			return true
		} else {
			p = p.next
		}
	}
	return false
}

func (l *LinkList) Insert(p int, e interface{}) bool {
	if p < 0 {
		return false
	} else if p == 0 {
		l.InsertHead(e)
		return true
	} else if p == l.size {
		l.Append(e)
		return true
	} else if p > l.size {
		return false
	} else {
		index := 0
		point := l.head
		for index = 0; index < p-1; index++ {
			point = point.next
		}
		l.InsertAfterNode(point, e)
		return true
	}
}

func (l *LinkList) DeleteNode(node *Node) {
	if l.exit(node) {
		if node == l.head {
			l.head = l.head.next
		} else if node == l.tail {
			p := l.head
			for p.next != l.tail {
				p = p.next
			}
			p.next = nil
			l.tail = p
		} else {
			p := l.head
			for p.next != node {
				p = p.next
			}
			p.next = node.next
		}
		l.size--
	}
}

func (l *LinkList) DeleteData(e interface{}) {
	p := l.Get(e)
	if p == nil {
		fmt.Println("is null")
	} else {
		l.DeleteNode(p)
	}
}
