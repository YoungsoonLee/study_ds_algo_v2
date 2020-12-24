package main

import (
	"fmt"
	"math"
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
