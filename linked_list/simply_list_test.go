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
	expected := []interface{}{2, 1}

	list.AddFirst(&node1)
	list.AddFirst(&node2)

	assert.Equal(expected, list.ToArray())
	assert.Equal(list.head, &node2)
	assert.Equal(list.tail, &node1)
}

func TestSimplyList_AddLast(t *testing.T) {
	assert := assert.New(t)
	list := SimplyList{}
	node1 := Node{1, nil}
	node2 := Node{2, &node1}
	expected := []interface{}{2, 1}

	list.AddFirst(&node2)
	list.AddLast(&node1)

	assert.Equal(expected, list.ToArray())
	assert.Equal(list.head, &node2)
	assert.Equal(list.tail, &node1)
}

func TestSimplyList_InsertAfter(t *testing.T) {
	assert := assert.New(t)
	list := SimplyList{}
	node1 := Node{1, nil}
	node2 := Node{2, nil}
	expected := []interface{}{1, 2}

	list.AddFirst(&node1)
	list.InsertAfter(&node1, &node2)

	assert.Equal(expected, list.ToArray())
	assert.Equal(list.head.data, node1.data)
	assert.Equal(list.tail.data, node2.data)
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

	assert.Equal(list.head, &node2);
}

func TestSimplyList_RemoveAfter(t *testing.T) {
	assert := assert.New(t)
	list := SimplyList{}
	node1 := Node{1, nil}
	node2 := Node{2, nil}
	node3 := Node{3, nil}
	list.AddFirst(&node1)
	list.InsertAfter(&node1, &node2)
	list.InsertAfter(&node2, &node3)
	expected := []interface{}{1, 2}
	list.RemoveAfter(&node2)

	assert.Equal(expected, list.ToArray())
}

func TestSimplyList_RemoveNodeHasData(t *testing.T) {
	assert := assert.New(t)
	list := SimplyList{}
	node1 := Node{1, nil}
	node2 := Node{2, nil}
	node3 := Node{3, nil}
	list.AddFirst(&node1)
	list.InsertAfter(&node1, &node2)
	list.InsertAfter(&node2, &node3)
	expected := []interface{}{1, 3}
	list.RemoveNodeHasData(2)

	assert.Equal(expected, list.ToArray())
}

func TestSimplyList_RemoveList(t *testing.T) {
	assert := assert.New(t)
	list := SimplyList{}
	node1 := Node{1, nil}
	node2 := Node{2, nil}
	node3 := Node{3, nil}
	list.AddFirst(&node1)
	list.InsertAfter(&node1, &node2)
	list.InsertAfter(&node2, &node3)

	list.RemoveList()

	assert.Nil(list.head)
	assert.Nil(list.tail)
}

func TestSimplyList_ToArray(t *testing.T) {
	assert := assert.New(t)
	list := SimplyList{}
	node1 := Node{1, nil}
	node2 := Node{2, nil}
	node3 := Node{3, nil}
	list.AddFirst(&node1)
	list.InsertAfter(&node1, &node2)
	list.InsertAfter(&node2, &node3)
	expected := []interface{}{1, 2, 3}

	assert.Equal(expected, list.ToArray())
}