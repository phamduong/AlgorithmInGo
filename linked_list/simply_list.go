package linked_list

import "github.com/phamduong/AlgorithmInGo/cleaner"

type Node struct {
	data interface{}
	next *Node
}

func (node Node) NewNode(data interface{}, next *Node) Node {
	return Node{data, next}
}

type SimplyList struct {
	head, tail *Node
}

func (list *SimplyList) AddFirst(node *Node) {
	if list.head == nil {
		list.head = node
		list.tail = list.head
	} else {
		node.next = list.head
		list.head = node
	}
}

func (list *SimplyList) AddLast(node *Node) {
	if list.head == nil {
		list.head = node
		list.tail = list.head
	} else {
		list.tail.next = node
		list.tail = node
	}
}

func (list *SimplyList) InsertAfter(node *Node, newNode *Node) {
	if node != nil {
		newNode.next = node.next
		node.next = newNode
		if list.tail == node {
			list.tail = newNode
		}
	} else {
		list.AddFirst(newNode)
	}
}

func (list SimplyList) Search(x int) *Node {
	var node *Node = list.head
	for node != nil && node.data != x {
		node = node.next
	}

	return node
}

func (list *SimplyList) RemoveHead() {
	if list.head != nil {
		var node *Node = list.head
		list.head = list.head.next
		cleaner.Clear(node)
		if list.head == nil {
			list.tail = nil
		}
	}
}

func (list *SimplyList) RemoveAfter(node *Node) {
	if node != nil {
		var temp *Node = node.next
		if temp != nil {
			if temp == list.tail {
				list.tail = node
			}
			node.next = temp.next
			cleaner.Clear(temp)
		} else {
			list.RemoveHead()
		}
	}
}

func (list *SimplyList) RemoveNodeHasData(x int) int {
	var p, q *Node = list.head, nil
	for p != nil {
		if p.data == x {
			break
		}
		q = p;
		p = p.next
	}

	if p == nil {
		return 0
	} else {
		list.RemoveAfter(q)
		return 1
	}
}

func (list *SimplyList) RemoveList() {
	var node *Node
	for list.head != nil {
		node = list.head
		list.head = list.head.next
		cleaner.Clear(node)
	}
	list.tail = nil
}

func (list *SimplyList) ToArray() []interface{} {
	var s []interface{}
	var node *Node = list.head

	for node != nil {
		s = append(s, node.data)
		node = node.next
	}

	return s
}