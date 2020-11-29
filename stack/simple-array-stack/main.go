package main

import (
	"errors"
	"fmt"
)

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

type pair struct {
	open  rune
	close rune
}

var pairs = []pair{
	{'(', ')'},
	{'[', ']'},
	{'{', '}'},
}

func isValid(s string) bool {
	stack := NewStack(1)
	for _, r := range s {
		for _, p := range pairs {
			temp, _ := stack.Peek()
			if r == p.open {
				stack.Push(r)
				break
			} else if r == p.close && stack.IsEmpty() {
				return false
			} else if r == p.close && temp == p.open {
				stack.Pop()
				break
			} else if r == p.close && temp != p.open {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func reverse(str string) string {
	result := []rune(str)
	var beg int
	end := len(result) - 1
	for beg < end {
		result[beg], result[end] = result[end], result[beg]
		beg = beg + 1
		end = end - 1
	}
	return string(result)
}

func isPalindrome(str string) bool {
	if reverse(str) == str {
		return true
	}
	return false
}

func isPalindrome2(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func isPalindromeStack(str string) bool {
	stack := NewStack(1)
	i, n := 0, len(str)
	for i < n/2 {
		stack.Push(str[i])
		i++
	}

	if n%2 == 1 {
		i++
	}

	for i < len(str) {
		poped, _ := stack.Pop()
		if stack.IsEmpty() || str[i] != poped {
			return false
		}
		i++
	}
	return true
}

func (s *Stack) reverseStack() {
	if s.IsEmpty() {
		return
	}

	data, _ := s.Pop()
	s.reverseStack()
	s.insertAtBottom(data)
}

func (s *Stack) insertAtBottom(data interface{}) {
	if s.IsEmpty() {
		s.Push(data)
		return
	}
	temp, _ := s.Pop()
	s.insertAtBottom(data)
	s.Push(temp)
}

func main() {
	s := NewStack(3)
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println(s.Size())
	fmt.Println(s.Peek())
	s.reverseStack()
	fmt.Println(s.Peek())
}
