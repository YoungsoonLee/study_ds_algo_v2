package main

import (
	"bytes"
	"fmt"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

type Queue struct {
	array    []interface{}
	front    int
	rear     int
	capacity int
	size     int
}

func New(capacity int) *Queue {
	return new(Queue).Init(capacity)
}

func (q *Queue) Init(capacity int) *Queue {
	q.array = make([]interface{}, capacity)
	q.front = -1
	q.rear = -1
	q.size = 0
	q.capacity = capacity
	return q
}

func (q *Queue) length() int {
	return q.size
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func (q *Queue) isFull() bool {
	return q.size == q.capacity
}

func (q *Queue) String() string {
	var result bytes.Buffer
	result.WriteByte('[')
	j := q.front
	for i := 0; i < q.size; i++ {
		result.WriteString(fmt.Sprintf("%v", q.array[j]))
		if i < q.size-1 {
			result.WriteByte(' ')
		}
		j = (j + 1) % q.capacity
	}
	result.WriteByte(']')
	return result.String()
}

func (q *Queue) Front() interface{} {
	return q.array[q.front]
}

func (q *Queue) Back() interface{} {
	return q.array[q.rear]
}

func (q *Queue) enQueue(v interface{}) {
	if q.isFull() {
		return
	}

	q.rear = (q.rear + 1) % q.capacity
	q.array[q.rear] = v
	if q.front == -1 {
		q.front = q.rear
	}
	q.size++
}

func (q *Queue) deQueue() interface{} {
	if q.isEmpty() {
		return MinInt
	}

	data := q.array[q.front]
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
		q.size = 0
	} else {
		q.front = (q.front + 1) % q.capacity
		q.size--
	}
	return data
}
