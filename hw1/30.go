package main

import (
	"errors"
)

type Queue struct {
	inStack  []int
	outStack []int
}

func (q *Queue) Enqueue(value int) {
	q.inStack = append(q.inStack, value)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.outStack) == 0 {
		if len(q.inStack) == 0 {
			return 0, errors.New("queue is empty")
		}

		for len(q.inStack) > 0 {
			last := len(q.inStack) - 1
			q.outStack = append(q.outStack, q.inStack[last])
			q.inStack = q.inStack[:last]
		}
	}

	last := len(q.outStack) - 1
	value := q.outStack[last]
	q.outStack = q.outStack[:last]

	return value, nil
}
