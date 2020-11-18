package singlyLinkedList

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
		return fmt.Errorf("insert: index out of bounds")
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

func kthFromEnd(head *ListNode, n int) *ListNode {
	first, second := head, head
	for ; n > 0; n-- {
		second = second.next
	}
	for ; second.next != nil; first, second = first.next, second.next {
	}
	first.next = first.next.next
	return first
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}
	}
	return false
}

func findLoopBeginning(head *ListNode) *ListNode {
	fast, slow := head, head
	loopExists := false
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			loopExists = true
			break
		}
	}

	if loopExists {
		slow = head
		for slow != fast {
			fast = fast.next
			slow = slow.next
		}
		return slow
	}
	return nil
}

func findLoopBeginning2(head *ListNode) int {
	fast, slow := head, head
	loopExists := false
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			loopExists = true
			break
		}
	}
	if loopExists {
		counter := 1
		fast = fast.next
		for slow != fast {
			fast = fast.next
			counter++
		}
		return counter
	}
	return 0
}

func (ll *LinkedList) sortedInsert(data int) {
	newNode := &ListNode{
		data: data,
	}
	if ll.head == nil || ll.head.data.(int) >= data {
		newNode.next = ll.head
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil && current.next.data.(int) < data {
		current = current.next
	}
	newNode.next = current.next
	current.next = newNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, current *ListNode

	for current = head; current != nil; current = current.next {
		current.next, prev, current = prev, current, current.next
	}
	return prev
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	h := reverse(head)
	head.next = nil
	return h
}

func reverse(current *ListNode) *ListNode {
	if current == nil {
		return nil
	}
	temp := reverse(current.next)
	if temp == nil {
		return current
	}
	current.next.next = current
	return temp
}

func getIntersectionNode(head1, head2 *ListNode) *ListNode {
	for head1 != nil {
		temp := head2
		for temp != nil {
			if temp == head1 {
				return head1
			}
			temp = temp.next
		}
		head1 = head1.next
	}
	return nil
}

func getIntersectionNode2(head1, head2 *ListNode) *ListNode {
	len1, len2 := findLen(head1), findLen(head2)
	if len1 > len2 {
		for ; len1 > len2; len1-- {
			head1 = head1.next
		}
	} else {
		for ; len2 > len1; len2-- {
			head2 = head2.next
		}
	}
	for head1 != head2 {
		head1, head2 = head1.next, head2.next
	}
	return head1
}

func findLen(head *ListNode) int {
	i := 0
	for ; head != nil; head = head.next {
		i++
	}
	return i
}

func middleNode(head *ListNode) *ListNode {
	i := findLen(head)
	count, target := 0, (i/2)+1
	for {
		count++
		if count == target {
			return head
		}
		head = head.next
	}
}

func middleNode2(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}

func printListInReverse(head *ListNode) {
	if head == nil {
		return
	}
	printListInReverse(head.next)
	fmt.Print(head.data)
}

func (ll *LinkedList) IsLengthEven() bool {
	current := ll.head
	for current != nil && current.next != nil {
		current = current.next.next
	}
	if current != nil {
		return false
	}
	return true
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.data.(int) < l2.data.(int) {
		l1.next = mergeTwoLists(l1.next, l2)
		return l1
	}
	l2.next = mergeTwoLists(l1, l2.next)
	return l2
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	for node := dummy; l1 != nil || l2 != nil; node = node.next {
		if l1 == nil {
			node.next = l2
			break
		} else if l2 == nil {
			node.next = l1
			break
		} else if l1.data.(int) < l2.data.(int) {
			node.next = l1
			l1 = l1.next
		} else {
			node.next = l2
			l2 = l2.next
		}
	}
	return dummy.next
}
