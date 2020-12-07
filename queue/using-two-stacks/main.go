package main

type Queue struct {
	stack1 []int
	stack2 []int
}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) EnQueue(data int) {
	q.stack1 = append(q.stack1, data)
}

func (q *Queue) DeQueue() int {
	if len(q.stack2) == 0 {
		for len(q.stack1) > 0 {
			item := q.stack1[len(q.stack1)-1]
			q.stack1 = q.stack1[:len(q.stack1)-1]
			q.stack2 = append(q.stack2, item)
		}
	}
	item := q.stack2[len(q.stack2)-1]
	q.stack2 = q.stack2[:len(q.stack2)-1]
	return item
}

func (q *Queue) Front() int {
	if len(q.stack2) < 1 {
		for len(q.stack1) > 0 {
			item := q.stack1[len(q.stack1)-1]
			q.stack1 = q.stack1[:len(q.stack1)-1]
			q.stack2 = append(q.stack2, item)
		}
	}
	return q.stack2[len(q.stack2)-1]
}

func (q *Queue) IsEmpty() bool {
	return len(q.stack1) == 0 && len(q.stack2) == 0
}
