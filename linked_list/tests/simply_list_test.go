package linked_list

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/phamduong/AlgorithmInGo/linked_list"
)

func TestAddFirst(t *testing.T) {
	assert := assert.New(t)
	list := linked_list.SimplyList{}
	node1 := linked_list.Node{1, nil}
	node2 := linked_list.Node{2, &node1}
	expected := []int{2, 1}

	list.AddFirst(node1)
	list.AddFirst(node2)

	assert.Equal(expected, list.ToArray())
	assert.Equal(list.Head, &node2)
}

func TestAddLast(t *testing.T) {
	assert := assert.New(t)
	list := linked_list.SimplyList{}
	node1 := linked_list.Node{1, nil}
	node2 := linked_list.Node{2, &node1}
	expected := []int{2, 1}

	list.AddFirst(node2)
	list.AddLast(node1)

	assert.Equal(expected, list.ToArray())
	assert.Equal(list.Head, &node2)
	assert.Equal(list.Tail, &node1)
}