package main

type Stack struct {
	top  *ListNode
	size int
}

type ListNode struct {
	data interface{}
	next *ListNode
}

func (s *Stack) length() int {
	return s.size
}

func (s *Stack) isEmpty() bool {
	return s.size == 0
}

func (s *Stack) isFull() bool {
	return false
}

func (s *Stack) push(data interface{}) {
	if !s.isFull() {
		//newNode := &ListNode{data, s.top}
		//s.top = newNode
		s.top = &ListNode{data, s.top}
		s.size++
	}
}

func (s *Stack) pop() (data interface{}) {
	if !s.isEmpty() {
		data, s.top = s.top.data, s.top.next
		s.size--
		return
	}
	return nil
}

func (s *Stack) peek() (data interface{}) {
	if !s.isEmpty() {
		data = s.top.data
		return
	}
	return nil
}
