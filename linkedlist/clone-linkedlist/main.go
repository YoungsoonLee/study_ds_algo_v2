package main

type ListNode struct {
	data   interface{}
	next   *ListNode
	random *ListNode
}

func clone(head *ListNode) *ListNode {
	var (
		result *ListNode               = &ListNode{}
		Y                              = result
		X      *ListNode               = head
		HT     map[*ListNode]*ListNode = make(map[*ListNode]*ListNode, 0)
	)

	for X != nil {
		Y.next = &ListNode{data: X.data}
		HT[X] = Y.next
		Y = Y.next
		X = X.next
	}
	X = head
	Y = result.next
	for X != nil {
		if n, found := HT[X.random]; found {
			Y.random = n
		}
		Y = Y.next
		X = X.next
	}
	return result.next
}
