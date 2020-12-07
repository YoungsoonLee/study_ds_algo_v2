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

func findMin(A []int, i, j int) int {
	min := A[i]
	for i <= j {
		if min > A[i] {
			min = A[i]
		}
		i++
	}
	return min
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for j, minimum_height := i, heights[i]; j < len(heights); j++ {
			minimum_height = findMin(heights, i, j)
			maxArea = findMax(maxArea, (j-i+1)*minimum_height)
		}
	}
	return maxArea
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleAreaStack(heights []int) int {
	i, max := 0, 0
	stack := make([]int, 0)

	for i < len(heights) {
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[pop]
			var w int
			if len(stack) == 0 {
				w = i
			} else {
				w = i - stack[len(stack)-1] - 1
			}
			max = findMax(max, h*w)
		}
	}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		h := heights[pop]
		var w int
		if len(stack) == 0 {
			w = i
		} else {
			w = i - stack[len(stack)-1] - 1
		}
		max = findMax(max, h*w)
	}

	return max
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
func pairWiseConsecutive(s *Stack) bool {
	auxStack := NewStack(s.Size())
	for !s.IsEmpty() {
		auxStack.Push(s.Peek())
		s.Pop()
	}

	result := true
	for auxStack.Size() > 1 {
		x, _ := auxStack.Peek()
		auxStack.Pop()
		y, _ := auxStack.Peek()
		auxStack.Pop()
		if abs(x-y) != 1 {
			result = false
		}

		s.Push(x)
		s.Push(y)
	}

	if auxStack.Size() == 1 {
		s.Push(auxStack.Peek())
	}
	return result

}
*/

func removeDuplicates(S string) string {
	stack := make([]byte, 0, len(S))
	for i := range S {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(int(14 / 10))
	/*
		s := NewStack(3)
		s.Push(10)
		s.Push(20)
		s.Push(30)
		fmt.Println(s.Size())
		fmt.Println(s.Peek())
		s.reverseStack()
		fmt.Println(s.Peek())
	*/
}
