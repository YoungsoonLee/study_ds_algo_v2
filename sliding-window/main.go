package main

type DeQueue struct {
	indexes []int
}

func (d *DeQueue) push(i int) {
	d.indexes = append(d.indexes, i)
}

func (d *DeQueue) getFirst() int {
	return d.indexes[0]
}

func (d *DeQueue) popFirst() {
	d.indexes = d.indexes[1:]
}

func (d *DeQueue) getLast() int {
	return d.indexes[len(d.indexes)-1]
}

func (d *DeQueue) popLast() {
	d.indexes = d.indexes[:len(d.indexes)-1]
}

func (d *DeQueue) empty() bool {
	return 0 == len(d.indexes)
}

func maxSlidingWindow(A []int, k int) []int {
	if len(A) < k || 0 == k {
		return make([]int, 0)
	} else if 1 == k {
		return A
	}

	var (
		// len(A)-k+1 window size.
		res = make([]int, len(A)-k+1)
		dq  = &DeQueue{}
	)

	for i := range A {
		if false == dq.empty() && (i-k == dq.getFirst()) {
			dq.popFirst()
		}

		for false == dq.empty() && A[dq.getLast()] < A[i] {
			dq.popLast()
		}

		dq.push(i)

		if i >= k-1 {
			res[i-k+1] = A[dq.getFirst()]
		}
	}
	return res
}
