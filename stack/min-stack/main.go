package main

import "math"

type MinStack struct {
	elementStack []int
	minimumStack []int
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func NewMinStack() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(data int) {
	s.elementStack = append(s.elementStack, data)
	if len(s.minimumStack) == 0 {
		s.minimumStack = append(s.minimumStack, data)
	} else {
		minimum := min(s.minimumStack[len(s.minimumStack)-1], data)
		s.minimumStack = append(s.minimumStack, minimum)
	}
}

func (s *MinStack) Pop() int {
	if len(s.elementStack) > 0 {
		poped := s.elementStack[len(s.elementStack)-1]
		s.elementStack = s.elementStack[:len(s.elementStack)-1]
		s.minimumStack = s.minimumStack[:len(s.minimumStack)-1]
		return poped
	} else {
		return math.MaxInt32
	}
}

func (s *MinStack) Peek() int {
	if len(s.elementStack) > 0 {
		return s.elementStack[len(s.elementStack)-1]
	} else {
		return 0
	}
}

func (s *MinStack) Size() int {
	return len(s.elementStack)
}

func (s *MinStack) GetMin() int {
	if len(s.minimumStack) > 0 {
		return s.minimumStack[len(s.minimumStack)-1]
	} else {
		return 0
	}
}

func (s *MinStack) IsEmpty() bool {
	return len(s.elementStack) == 0
}

func (s *MinStack) Clear() {
	s.elementStack = nil
	s.minimumStack = nil
}
