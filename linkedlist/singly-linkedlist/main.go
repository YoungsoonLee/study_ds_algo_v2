package singlyLinkedList

import (
	"fmt"
)

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

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		l1 := lists[0]
		l2 := lists[1]
		lists = lists[2:]
		merged := mergeTwoLists(l1, l2)
		lists = append(lists, merged)
	}
	return lists[0]
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}

	slow, fast := head, head
	for fast.next != nil && fast.next.next != nil {
		slow, fast = slow.next, fast.next.next
	}

	firstTail := slow
	slow = slow.next
	firstTail.next = nil

	first, second := sortList(head), sortList(slow)
	return merge(first, second)
}

func merge(head1 *ListNode, head2 *ListNode) *ListNode {
	curHead := &ListNode{}
	tmpHead := curHead

	for head1 != nil && head2 != nil {
		if head1.data.(int) < head2.data.(int) {
			curHead.next = head1
			head1 = head1.next
			curHead = curHead.next
		} else {
			curHead.next = head2
			head2 = head2.next
			curHead = curHead.next
		}
	}

	if head1 != nil {
		curHead.next = head1
	} else if head2 != nil {
		curHead.next = head2
	}
	return tmpHead.next
}

func splitList(head *ListNode) (head1 *ListNode, head2 *ListNode) {
	var slow, fast *ListNode
	if head == nil || head.next == nil {
		head1 = head
		head2 = nil
	} else {
		slow = head
		fast = head.next
		for fast != nil {
			slow = slow.next
			fast = fast.next.next
		}
		head1 = head
		head2 = slow.next
		slow.next = nil
	}
	return head1, head2
}

func reverseBlockOfKNodes(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	length := 0
	node := head
	for node != nil {
		length++
		node = node.next
	}

	result := ListNode{0, head}
	previous := &result
	for step := 0; step+k <= length; step = step + k {
		tail := previous.next
		nextNode := tail.next
		for i := 1; i < k; i++ {
			tail.next = nextNode.next
			nextNode.next = previous.next
			previous.next = nextNode
			nextNode = tail.next
		}
		previous = tail
	}
	return result.next
}

func (ll *LinkedList) modularNodeFromBegin(k int) *ListNode {
	if k < 0 {
		return nil
	}
	i := 1

	current, modularNode := ll.head, ll.head
	for ; current != nil; current = current.next {
		if i%k == 0 {
			modularNode = ll.head
		} else {
			modularNode = modularNode.next
		}
		current = current.next
		i++
	}
	return modularNode
}

func (ll *LinkedList) modularNodeFromEnd(k int) *ListNode {
	if k <= 0 {
		return nil
	}
	current, modularNode := ll.head, ll.head
	i := 0
	for i := 0; i < k; i++ {
		if current != nil {
			current = current.next
		} else {
			break
		}
	}

	for current != nil {
		modularNode = modularNode.next
		current = current.next
		i++
	}

	j := k - (i % k)
	for j > 0 && modularNode != nil {
		modularNode = modularNode.next
		j--
	}
	return modularNode
}

func (ll *LinkedList) fractionalNode(k int) *ListNode {
	if k <= 0 {
		return nil
	}

	i := 0
	current := ll.head
	var fractionalNode *ListNode
	for ; current != nil; current = current.next {
		if i%k == 0 {
			if fractionalNode == nil {
				fractionalNode = ll.head
			} else {
				fractionalNode = fractionalNode.next
			}
		}
		i++
	}
	return fractionalNode
}

func (ll *LinkedList) sqrtNode() *ListNode {
	current := ll.head
	var sqrtN *ListNode
	for i, j := 1, 1; current != nil; current = current.next {
		if i == j*j {
			if sqrtN == nil {
				sqrtN = ll.head
			} else {
				sqrtN = sqrtN.next
			}
			j++
		}
		i++
	}
	return sqrtN
}

func mergeTwoListsWithOrder(head1 *ListNode, head2 *ListNode) *ListNode {
	h := ListNode{}
	l := &h

	for head1 != nil && head2 != nil {
		if head1.data.(int) <= head2.data(int) {
			l.next = head1
			head1 = head1.next
		} else {
			l.next = head2
			head2 = head2.next
		}
		l = l.next
	}

	if head1 == nil {
		l.next = head2
	}
	if head2 == nil {
		l.next = head1
	}
	return h.next
}

func mergeTwoListsWithOrder2(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	if head1.data.(int) < head2.data.(int) {
		head1.next = mergeTwoListsWithOrder2(head1.next, head2)
		return head1
	}
	head2.next = mergeTwoListsWithOrder2(head1, head2.next)
	return head2
}

func segregateEvenOdds(head *ListNode) *ListNode {
	var evensHead, evenEnd, oddsHead, oddEnd *ListNode
	evensHead, evenEnd, oddsHead, oddEnd = nil, nil, nil, nil

	currNode := head
	for currNode != nil {
		val := currNode.data.(int)
		if val%2 == 0 {
			if evensHead == nil {
				evensHead = currNode
				evenEnd = evensHead
			} else {
				evenEnd.next = currNode
				evenEnd = evenEnd.next
			}
		} else {
			if oddsHead == nil {
				oddsHead = currNode
				oddEnd = oddsHead
			} else {
				oddEnd.next = currNode
				oddEnd = oddEnd.next
			}
		}
		currNode = currNode.next
	}

	if oddsHead == nil || evensHead == nil {
		return head
	}

	evenEnd.next = oddsHead
	oddEnd.next = nil
	return evensHead
}

func reorderList(head *ListNode) {
	if head == nil || head.next == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow, fast = slow.next, fast.next.next
	}

	var prev *ListNode
	for slow != nil {
		slow.next = prev
		prev = slow
		slow = slow.next
	}

	first := head
	for prev.next != nil {
		first.next, first = prev, first.next
		prev.next, prev = first, prev.next
	}
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.next == nil {
		return head
	}

	oddsHead := head
	evenHead := head.next
	for current, temp := head, head.next; temp != nil; current, temp = temp, temp.next {
		current.next = temp.next
	}

	oddsTail := oddsHead
	for ; oddsTail.next != nil; oddsTail = oddsTail.next {
	}

	oddsTail.next = evenHead
	return oddsHead
}

func reversePairs(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}

	result := head.next
	head.next = swapPairs(head.next.next)
	result.next = head
	return result
}

func reversePairs2(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}

	result := head.next
	var previousNode *ListNode
	for head != nil && head.next != nil {
		nextNode := head.next
		head.next = nextNode.next
		nextNode.next = head
		if previousNode != nil {
			previousNode.next = nextNode
		}
		previousNode = head
		head = head.next
	}
	return result
}

func partition(head *ListNode, X int) *ListNode {
	lesser, greater := &ListNode{}, &ListNode{}
	lesserHead, greaterHead := lesser, greater
	for head != nil {
		if head.data.(int) < X {
			lesser.next = head
			lesser = lesser.next
		} else {
			greater.next = head
			greater = greater.next
		}
		head = head.next
	}
	lesser.next, greater.next = nil, nil
	if greaterHead.next != nil {
		lesser.next = greaterHead.next
	}
	return lesserHead.next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, result := 0, new(ListNode)
	for node := result; l1 != nil || l2 != nil || carry > 0; node = node.next {
		if l1 != nil {
			carry += l1.data.(int)
			l1 = l1.next
		}
		if l2 != nil {
			carry += l2.data.(int)
			l2 = l2.next
		}
		node.next = &ListNode{carry % 10, nil}
		carry /= 10
	}
	return result.next
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	result := &ListNode{next: head}
	current := head.next
	head.next = nil

	for current != nil {
		pre := result
		target := result.next
		for target != nil && current.data.(int) > target.data.(int) {
			target = target.next
			pre = pre.next
		}
		temp := current
		current = current.next
		temp.next = target
		pre.next = temp
	}
	return result.next
}

func intersection(list1, list2 *ListNode) *ListNode {
	list := LinkedList{}
	for list1 != nil && list2 != nil {
		if list1.data.(int) == list2.data.(int) {
			list.InsertBeginning(list1.data)
			list1 = list1.next
			list2 = list2.next
		} else if list1.data.(int) > list2.data.(int) {
			list2 = list2.next
		} else {
			list1 = list1.next
		}
	}
	return list.head
}

func intersection2(list1, list2 *ListNode) *ListNode {
	headList := LinkedList{}
	tailList := LinkedList{}
	head, tail := headList.head, tailList.head
	for list1 != nil && list2 != nil {
		if list1.data == list2.data {
			if head == nil {
				headList.InsertBeginning(list1.data)
				tailList = headList
			} else {
				tailList.InsertBeginning(list1.data)
				tail = tail.next
			}
			list1 = list1.next
			list2 = list2.next
		} else if list1.data.(int) > list2.data.(int) {
			list2 = list2.next
		} else {
			list1 = list1.next
		}
	}
	return headList.head
}

func intersection3(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	list := LinkedList{}
	list.head = dummy
	for list1 != nil && list2 != nil {
		if list1.data == list2.data {
			list.InsertBeginning(list1.data)
			list1 = list1.next
			list2 = list2.next
		} else if list1.data.(int) > list2.data.(int) {
			list2 = list2.next
		} else {
			list1 = list1.next
		}
	}
	return list.head
}

func removeDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		if current.next != nil && current.data == current.next.data {
			current.next = current.next.next
			continue
		}
		current = current.next
	}
	return head
}

func alternatingSplit(head *ListNode) (head1, head2 *ListNode) {
	var a *ListNode = nil
	var b *ListNode = nil
	current := head
	for current != nil {
		newNode := current
		current = newNode.next
		newNode.next = a
		a = newNode
		if current != nil {
			newNode := current
			current = newNode.next
			newNode.next = b
			b = newNode
		}
	}
	head1, head2 = a, b
	return head1, head2
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	type SumNode struct {
		Node *ListNode
		Sum  int
	}
	acc, sum, dummy := 0, make(map[int]SumNode), &ListNode{next: head}
	for curr := head; curr != nil; curr = curr.next {
		acc += curr.data.(int)
		if p, ok := sum[acc]; ok {
			for node, subSum := p.Node.next, p.Sum; node != curr; node = node.next {
				subSum += node.data.(int)
				delete(sum, subSum)
			}
			p.Node.next = curr.next
		} else {
			sum[acc] = SumNode{curr, acc}
		}
	}
	return dummy.next
}
