package main

type DLLNode struct {
	data interface{}
	prev *DLLNode
	next *DLLNode
}

type DLL struct {
	head *DLLNode
	tail *DLLNode
	size int
}

func (dll *DLL) CheckIffEmptyAndAdd(newNode *DLLNode) bool {
	if dll.size == 0 {
		dll.head = newNode
		dll.tail = newNode
		dll.size++
		return true
	}
	return false
}

func (dll *DLL) InsertBeginning(data int) {
	newNode := &DLLNode{data: data, prev: nil, next: nil}
	if !(dll.CheckIffEmptyAndAdd(newNode)) {
		head := dll.head
		newNode.next = head
		newNode.prev = nil

		head.prev = newNode
		dll.head = newNode
		dll.size++
		return
	}
	return
}

func (dll *DLL) InsertEnd(data int) {
	newNode := &DLLNode{
		data: data,
		prev: nil,
		next: nil,
	}

	if !(dll.CheckIffEmptyAndAdd(newNode)) {
		head := dll.head
		for i := 0; i < dll.size; i++ {
			if head.next == nil {
				newNode.prev = head
				newNode.next = nil

				head.next = newNode

				dll.tail = newNode
				dll.size++
				break
			}
			head = head.next
		}
	}
	return
}

func (dll *DLL) Insert(data interface{}, loc int) {
	newNode := &DLLNode{
		data: data,
		prev: nil,
		next: nil,
	}

	if !(dll.CheckIffEmptyAndAdd(newNode)) {
		head := dll.head
		for i := 1; i < dll.size; i++ {
			if i == loc {
				newNode.prev = head.prev
				newNode.next = head

				head.prev.next = newNode
				head.prev = newNode
				dll.size++
				return
			}
			head = head.next
		}
	}
	return
}

func (dll *DLL) CheckIfEmpty() bool {
	return dll.size == 0
}

func (dll *DLL) DeleteFirst() int {
	if !(dll.CheckIfEmpty()) {
		head := dll.head
		if head.prev == nil {
			deletedNode := head.data
			dll.head = head.next
			dll.head.prev = nil
			dll.size--

			return deletedNode.(int)
		}
	}
	return -1
}

func (dll *DLL) DeleteLast() int {
	if !(dll.CheckIfEmpty()) {
		head := dll.head
		for {
			if head.next == nil {
				break
			}
			head = head.next
		}

		dll.tail = head.prev
		dll.tail.next = nil
		dll.size--
		return head.data.(int)
	}
	return -1
}
