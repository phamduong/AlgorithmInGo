package stack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStack_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	stack := Stack{nil, 0}

	assert.True(stack.IsEmpty())
}

func TestStack_Pop(t *testing.T) {
	assert := assert.New(t)
	stack := new(Stack)

	stack.Push("test")
	stack.Push("test1")

	assert.Equal(2, stack.Len())
	assert.Equal("test1", stack.Pop().(string))
	assert.Equal("test", stack.Pop().(string))
	assert.Equal(0, stack.Len())
}

func TestStack_Push(t *testing.T) {
	assert := assert.New(t)
	stack := new(Stack)

	stack.Push("test")

	assert.Equal(1, stack.Len())
	assert.Equal("test", stack.Pop().(string))
}

func TestStack_Top(t *testing.T) {
	assert := assert.New(t)
	stack := new(Stack)

	stack.Push("test")

	assert.Equal(1, stack.Len())
	assert.Equal("test", stack.Top().(string))
	assert.Equal(1, stack.Len())
}