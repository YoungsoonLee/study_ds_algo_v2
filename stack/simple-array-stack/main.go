package main

import "errors"

type Stack struct {
	top      int
	capacity uint
	array    []interface{}
}

func (stack *Stack) Init(capacity uint) *Stack {
	stack.top = -1
	stack.capacity = capacity
	stack.array = make([]interface{}, capacity)
	return stack
}

func NewStack(capacity uint) *Stack {
	return new(Stack).Init(capacity)
}

func (stack *Stack) IsFull() bool {
	return stack.top == int(stack.capacity)-1
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

func (stack *Stack) Size() uint {
	return uint(stack.top + 1)
}

func (stack *Stack) Push(data interface{}) error {
	if stack.IsFull() {
		return errors.New("stack is full")
	}
	stack.top++
	stack.array[stack.top] = data
	return nil
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	temp := stack.array[stack.top]
	stack.top--
	return temp, nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return stack.array[stack.top], nil
}

// Drain. remove all elemments
func (stack *Stack) Drain() {
	stack.array = nil
	stack.top = -1
}
