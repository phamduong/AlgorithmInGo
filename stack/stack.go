package stack

type Element struct {
	value interface{}
	next  *Element
}

type Stack struct {
	top  *Element
	size int
}

func (stack Stack) NewStack(top Element, size int) Stack {
	return Stack{&top, size}
}

func (stack Stack) IsEmpty() bool {
	return stack.size == 0
}

func (stack Stack) Len() int {
	return stack.size
}

func (stack Stack) Pop() (value interface{}) {
	if stack.size > 0 {
		value, stack.top = stack.top.value, stack.top.next
		stack.size--
		return
	}
	return nil
}

func (stack Stack) Top() (value interface{}) {
	if stack.size > 0 {
		return stack.top.value
	}
	return nil
}

func (stack Stack) Push(value interface{}) {
	stack.top = &Element{value, stack.top}
	stack.size++
}