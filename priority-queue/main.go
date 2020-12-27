package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Item interface {
	Less(item Item) bool
}

type Heap struct {
	size int
	data []Item
}

func New() *Heap {
	return &Heap{}
}

func parent(i int) int {
	return int(math.Floor(float64(i-1) / 2.0))
}

func leftChild(parent int) int {
	return (2 * parent) + 1
}

func rightChile(parent int) int {
	return (2 * parent) + 2
}

func getMinimum(h *Heap) (Item, error) {
	if h.size == 0 {
		return nil, fmt.Errorf("Unable to get element from empty heap")
	}
	return h.data[0], nil
}

// heapify when inserting.
func (h *Heap) percolateUp() {
	idx := h.size - 1

	if idx <= 0 {
		return
	}

	for {
		p := parent(idx)
		if p < 0 || h.data[p].Less(h.data[idx]) { // minheap
			break
		}
		swap(h, p, idx)
		idx = p
	}
}

func swap(h *Heap, i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// after deleting.
func (h *Heap) percolateDown(i int) {
	p := i
	for {
		l := leftChild(p)
		r := rightChile(p)

		s := p
		if l < h.size && h.data[l].Less(h.data[s]) {
			s = l
		}

		if r < h.size && h.data[r].Less(h.data[s]) {
			s = r
		}

		if s == p {
			break
		}

		swap(h, p, s)
		p = s

	}
}

func (h *Heap) Extract() (Item, error) {
	n := h.size
	if n == 0 {
		return nil, fmt.Errorf("Unable to extract from empty heap")
	}

	m := h.data[0]
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]
	h.size--

	if h.size > 0 {
		h.percolateDown(0)
	} else {
		h.data = nil
	}
	return m, nil
}

func (h *Heap) Insert(item Item) {
	if h.size == 0 {
		h.data = make([]Item, 1)
		h.data[0] = item
	} else {
		h.data = append(h.data, item)
	}
	h.size++
	h.percolateUp()
}

func Heapify(items []Item) *Heap {
	h := New()
	n := len(items)
	h.data = make([]Item, n)
	copy(h.data, items)
	h.size = len(items)
	i := int(n / 2)
	for i >= 0 {
		h.percolateDown(i)
		i--
	}
	return h
}

type Int int

func (a Int) Less(b Item) bool {
	val, ok := b.(Int)
	return ok && a <= val
}

func verifyHeap(h *Heap) bool {
	queue := make([]Int, 1)
	queue[0] = 0
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		l := leftChild(int(p))
		r := rightChile(int(p))
		if l < h.size {
			if !h.data[p].Less(h.data[l]) {
				return false
			}
			queue = append(queue, Int(l))
		}
		if r < h.size {
			if !h.data[p].Less(h.data[r]) {
				return false
			}
			queue = append(queue, Int(r))
		}
	}
	return true
}

func verifyStrictlyIncreasing(h *Heap) (bool, []Item) {
	prev, _ := h.Extract()
	order := []Item{prev}
	for h.size > 0 {
		curr, _ := h.Extract()
		order = append(order, curr)
		if curr.Less(prev) {
			return false, order
		}
		prev = curr
		order = append(order, prev)
	}
	return true, order
}

func randomPerm(n int) []Item {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	ints := r.Perm(n)
	items := make([]Item, n)
	for idx, item := range ints {
		items[idx] = Int(item)
	}
	return items
}

func HeapSort(data []Item) []Item {
	hp := Heapify(data)
	size := len(hp.data)
	for i := size - 1; i > 0; i-- {
		swap(hp, 0, i)
		hp.size--
		hp.percolateDown(0)
	}
	hp.size = size
	return hp.data
}

func findMaxInMinHeap(h *Heap) Int {
	max := Int(-1)
	for i := (h.size + 1) / 2; i < h.size; i++ {
		if h.data[i].(Int) > max {
			max = h.data[i].(Int)
		}
	}
	return max
}

func delete(h *Heap, i int) int {
	if i > h.size {
		fmt.Println("Wrong position")
		return -1
	}

	key := h.data[i].(Int)
	h.data[i] = h.data[h.size-1]
	h.size--
	h.percolateDown(i)
	return key
}

func main() {
	items := randomPerm(20)
	hp := New()
	for _, item := range items {
		fmt.Println("Inserting an element into Heap: ", hp.data)
		hp.Insert(item)
	}
	if !verifyHeap(hp) {
		fmt.Println("invalid Heap: ", hp.data)
		return
	}
	if ok, order := verifyStrictlyIncreasing(hp); !ok {
		fmt.Println("invalid Heap extraction order: ", order)
	}
}
