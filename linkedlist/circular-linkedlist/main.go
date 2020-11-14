package circularLinkedList

import (
	"fmt"
)

type CLLNode struct {
	data int
	next *CLLNode
}

type CLL struct {
	head *CLLNode
	size int
}

func (cll *CLL) Length() int {
	current := cll.head
	count := 1
	if current == nil {
		return 0
	}
	current = current.next
	for current != cll.head {
		current = current.next
		count++
	}
	return count
}

func (cll *CLL) Display() {
	head := cll.head
	for i := 0; i < cll.size; i++ {
		fmt.Print(head.data)
		fmt.Print("-->")
		head = head.next
	}
	fmt.Println()
}

func (cll *CLL) CheckIfEmptyAndAdd(data int) bool {
	newNode := &CLLNode{
		data: data,
		next: nil,
	}

	if cll.size == 0 {
		cll.head = newNode
		cll.head.next = newNode
		cll.size++
		return true
	}
	return false
}

func (cll *CLL) InsertBeginning(data int) {
	if !(cll.CheckIfEmptyAndAdd(data)) {
		newNode := &CLLNode{
			data: data,
			next: nil,
		}

		current := cll.head
		newNode.next = current
		for {
			if current.next == cll.head {
				break
			}
			current = current.next
		}
		current.next = newNode
		cll.head = newNode
		cll.size++
	}
}

func (cll *CLL) InsertEnd(data int) {
	if !(cll.CheckIfEmptyAndAdd(data)) {
		newNode := &CLLNode{
			data: data,
			next: nil,
		}

		current := cll.head
		for {
			if current.next == cll.head {
				break
			}
			current = current.next
		}
		current.next = newNode
		newNode.next = cll.head
		cll.size++

	}
}

func (cll *CLL) Insert(data int, pos int) {
	if !(cll.CheckIfEmptyAndAdd(data)) {
		current := cll.head
		count := 1

		if pos == 1 {
			cll.InsertBeginning(data)
			return
		}

		newNode := &CLLNode{
			data: data,
			next: nil,
		}
		for {
			if current.next == nil && pos-1 > count {
				break
			}
			if count == pos-1 {
				newNode.next = current.next
				current.next = newNode
				cll.size++
				break
			}
			current = current.next
			count++
		}
	}
}

func (cll *CLL) CheckIfEmpty() bool {
	return cll.size == 0
}

func (cll *CLL) DeleteBeginning() int {
	if !(cll.CheckIfEmpty()) {
		current := cll.head
		deletedElem := current.data
		if cll.size == 1 {
			cll.head = nil
			cll.size--
			return deletedElem
		}
		prevStart := cll.head
		cll.head = current.next
		for {
			if current.next == prevStart {
				break
			}
			current = current.next
		}
		current.next = cll.head
		cll.size--
		return deletedElem
	}
	return -1
}

func (cll *CLL) DeleteEnd() int {
	if !(cll.CheckIfEmpty()) {
		current := cll.head
		deletedEle := current.data
		if cll.size == 1 {
			deletedEle = cll.DeleteBeginning()
			return deletedEle
		}

		for {
			if current.next.next == cll.head {
				deletedEle = current.next.data
				break
			}
			current = current.next
		}

		current.next = cll.head
		cll.size--
		return deletedEle
	}
	return -1
}

func (cll *CLL) Delete(pos int) int {
	if !(cll.CheckIfEmpty()) {
		currenct := cll.head
		deletedEle := currenct.data
		if cll.size == 1 {
			deletedEle = cll.DeleteBeginning()
			return deletedEle
		}
		if cll.size == pos {
			deletedEle = cll.DeleteEnd()
			return deletedEle
		}
		count := 1
		for {
			if count == pos-1 {
				deletedEle = currenct.next.data
				break
			}
			currenct = currenct.next
		}
		currenct.next = currenct.next.next
		cll.size--
		return deletedEle
	}
	return -1
}
