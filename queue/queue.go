package queue

type Element struct {
	value interface{}
	next  *Element
}

type Queue struct {
	head, tail *Element
}

func (queue Queue) NewQueue(head, tail Element) Queue {
	return Queue{&head, &tail}
}

func (queue Queue) NewEmptyQueue() Queue {
	return Queue{}
}

func (queue Queue) EnQueue(value interface{}) {
	ele := &Element{value, nil}
	queue.tail.next, queue.tail = ele, ele
}

func (queue Queue) DeQueue() *Element {
	if queue.head != nil {
		ele := queue.head
		queue.head = queue.head.next
		return ele
	}
	return nil
}

func (queue Queue) IsEmpty() bool {
	return queue.head == nil
}

func (queue Queue) Front() *Element {
	return queue.head
}