package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimplyList_AddFirst(t *testing.T) {
	assert := assert.New(t)
	list := SimplyList{}
	node1 := Node{1, nil}
	node2 := Node{2, &node1}
	expected := []int{2, 1}

	list.AddFirst(&node1)
	list.AddFirst(&node2)

	assert.Equal(expected, list.ToArray())
	assert.Equal(list.Head, &node2)
	assert.Equal(list.Tail, &node1)
}

func TestSimplyList_AddLast(t *testing.T) {
	assert := assert.New(t)
	list := SimplyList{}
	node1 := Node{1, nil}
	node2 := Node{2, &node1}
	expected := []int{2, 1}

	list.AddFirst(&node2)
	list.AddLast(&node1)

	assert.Equal(expected, list.ToArray())
	assert.Equal(list.Head, &node2)
	assert.Equal(list.Tail, &node1)
}

func TestSimplyList_InsertAfter(t *testing.T) {
	assert := assert.New(t)
	list := SimplyList{}
	node1 := Node{1, nil}
	node2 := Node{2, nil}
	expected := []int{1, 2}

	list.AddFirst(&node1)
	list.InsertAfter(&node1, &node2)

	assert.Equal(expected, list.ToArray())
	assert.Equal(list.Head.Data, node1.Data)
	assert.Equal(list.Tail.Data, node2.Data)
}

func TestSimplyList_Search(t *testing.T) {
	assert := assert.New(t)
	list := SimplyList{}
	node1 := Node{1, nil}
	node2 := Node{2, nil}

	list.AddFirst(&node1)
	list.InsertAfter(&node1, &node2)

	assert.Equal(&node1, list.Search(1))
	assert.Equal(&node2, list.Search(2))
	assert.Nil(list.Search(3))
}

func TestSimplyList_RemoveHead(t *testing.T) {
	assert := assert.New(t)
	list := SimplyList{}
	node1 := Node{1, nil}
	node2 := Node{2, nil}
	list.AddFirst(&node1)
	list.InsertAfter(&node1, &node2)

	list.RemoveHead()

	assert.Equal(list.Head, &node2);
}