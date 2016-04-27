package linked_list

import "github.com/phamduong/AlgorithmInGo/cleaner"

type SimplyList struct {
	Head, Tail *Node
}

func (list *SimplyList) AddFirst(node Node) {
	if list.Head == nil {
		list.Head = &node
		list.Tail = list.Head
	} else {
		node.Next = list.Head
		list.Head = &node
	}
}

func (list *SimplyList) AddLast(node Node) {
	if list.Head == nil {
		list.Head = &node
		list.Tail = list.Head
	} else {
		list.Tail.Next = &node
		list.Tail = &node
	}
}

func (list *SimplyList) InsertAfter(node Node, newNode Node) {
	if &node != nil {
		newNode.Next = node.Next
		node.Next = &newNode
		if *list.Tail == node {
			list.Tail = &newNode
		}
	} else {
		list.AddFirst(newNode)
	}
}

func (list SimplyList) Search(x int) Node {
	var node *Node = list.Head
	for node != nil && node.Data != x {
		node = node.Next
	}

	return *node
}

func (list *SimplyList) RemoveHead() {
	if list.Head != nil {
		var node *Node = list.Head
		list.Head = list.Head.Next
		cleaner.Clear(*node)
		if list.Head == nil {
			list.Tail = nil
		}
	}
}

func (list *SimplyList) RemoveAfter(node Node) {
	if &node != nil {
		var temp *Node = node.Next
		if temp != nil {
			if temp == list.Tail {
				list.Tail = &node
			}
			node.Next = temp.Next
			cleaner.Clear(*temp)
		} else {
			list.RemoveHead()
		}
	}
}

func (list *SimplyList) RemoveNodeHasData(x int) int {
	var p, q *Node = list.Head, nil
	for p != nil {
		if p.Data == x {
			break
		}
		q = p;
		p = p.Next
	}

	if p == nil {
		return 0
	} else {
		list.RemoveAfter(*q)
		return 1
	}
}

func (list *SimplyList) RemoveList() {
	var node *Node
	for list.Head != nil {
		node = list.Head
		list.Head = list.Head.Next
		cleaner.Clear(node)
	}
	list.Tail = nil
}

func (list *SimplyList) ToArray() []int {
	var s []int
	var node *Node = list.Head

	for node != nil {
		s = append(s, node.Data)
		node = node.Next
	}

	return s
}