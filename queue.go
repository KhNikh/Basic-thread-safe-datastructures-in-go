package main

import (
	"errors"
	"fmt"
)

type queue struct {
	queue []interface{}
	size  int
}

func (q *queue) isEmpty() bool {
	return len(q.queue) == 0
}

func (q *queue) isFull() bool {
	return len(q.queue) == q.size
}

func (q *queue) front() (interface{}, error) {
	if q.isEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.queue[0], nil
}

func (q *queue) enqueue(val interface{}) error {
	if q.isFull() {
		return errors.New("queue is full")
	}
	q.queue = append(q.queue, val)
	return nil
}

func (q *queue) dequeue() error {
	if q.isEmpty() {
		return errors.New("queue is empty")
	}
	q.queue = q.queue[1:]
	return nil
}

func main() {
	que := queue{
		size: 10,
	}
	que.enqueue(10)
	front, err := que.front()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(front)
	}
	que.dequeue()

	front, err = que.front()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(front)
	}
}
