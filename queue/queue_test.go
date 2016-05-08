package queue

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestQueue_DeQueue(t *testing.T) {
	assert := assert.New(t)
	queue := Queue{}

	queue.EnQueue("element1")
	queue.EnQueue("element2")

	assert.Equal("element1", queue.DeQueue().value)
	assert.Equal("element2", queue.DeQueue().value)
	assert.Nil(queue.DeQueue())
}

func TestQueue_EnQueue(t *testing.T) {
	assert := assert.New(t)
	queue := Queue{}

	queue.EnQueue("element1")
	queue.EnQueue("element2")

	assert.Equal("element1", queue.Front().value)
	assert.Equal("element2", queue.tail.value)
}

func TestQueue_Front(t *testing.T) {
	assert := assert.New(t)
	queue := Queue{}

	queue.EnQueue("element1")
	assert.Equal("element1", queue.Front().value)
}

func TestQueue_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	queue := Queue{}

	assert.True(queue.IsEmpty())
}

func TestQueue_IsEmpty2(t *testing.T) {
	assert := assert.New(t)
	queue := Queue{}

	queue.EnQueue("element1")
	assert.False(queue.IsEmpty())
}