package main

import "fmt"

type ListNode struct {
	data interface{}
	next *ListNode
}

func NewListNode(data int) *ListNode {
	temp := &ListNode{}
	temp.next = temp
	temp.data = data
	return temp
}

func getJosephusPosition(m, n int) {
	head := NewListNode(1)
	prev := head
	for i := 2; i <= n; i++ {
		prev.next = NewListNode(i)
		prev = prev.next
	}
	prev.next = head

	ptr1, ptr2 := head, head
	for ptr1.next != ptr1 {
		count := 1
		for count != m {
			ptr2 = ptr1
			ptr1 = ptr1.next
			count++
		}
		//removve
		ptr2.next = ptr1.next
		ptr1 = ptr2.next
	}
	fmt.Println("Last person left standing ", "(Josephus Position) is ", ptr1.data)

}

func main() {
	n, m := 14, 2
	getJosephusPosition(m, n)
}
