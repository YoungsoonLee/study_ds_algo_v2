package main

import "fmt"

type ListNode struct {
	data interface{}
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func (ll *LinkedList) Display() error {
	if ll.head == nil {
		return fmt.Errorf("display: List is empty")
	}
	current := ll.head
	for ll.head != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println()
	return nil
}

func (ll *LinkedList) Length() int {
	return ll.size
}

func (ll *LinkedList) InsertBeginning(data interface{}) {
	node := &ListNode{data: data}
	if ll.head == nil {
		ll.head = node
	} else {
		node.next = ll.head
		ll.head = node
	}
	ll.size++
	return
}

func (ll *LinkedList) InsertEnd(data interface{}) {
	node := &ListNode{data: data}
	if ll.head == nil {
		ll.head = node
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}

	ll.size++
	return
}

func (ll *LinkedList) Insert(position int, data interface{}) error {
	node := &ListNode{data, nil}
	if position < 1 || position > ll.size+1 {
		return fmt.Errorf("innsert: index out of bounds")
	}

	var prev, current *ListNode
	prev = nil
	current = ll.head

	for position > 1 {
		prev = current
		current = current.next
		position--
	}

	if prev != nil {
		prev.next = node
		node.next = current
	} else {
		node.next = current
		ll.head = node
	}

	ll.size++
	return nil

}

func (ll *LinkedList) DeleteFirst() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}

	result := ll.head.data
	ll.head = ll.head.next
	ll.size--

	return result, nil
}

func (ll *LinkedList) DeleteLast() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deleteFront: List is empty")
	}

	var prev *ListNode
	prev = nil

	current := ll.head
	for current.next != nil {
		prev = current
		current = current.next
	}

	if prev != nil {
		prev.next = nil
	} else {
		ll.head = nil
	}

	ll.size--
	return current.data, nil

}

func (ll *LinkedList) Delete(position int) (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("deletePosition: List is empty")
	}

	var prev, current *ListNode
	prev = nil
	current = ll.head

	pos := 0
	if position == 1 {
		ll.head = ll.head.next
	} else {
		for pos != position-1 {
			prev = current
			current = current.next
			pos++
		}

		if current != nil {
			prev.next = current.next
		}
	}
	ll.size--
	return current.data, nil
}
