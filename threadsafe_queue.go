package main

import (
	"errors"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type safeQueue struct {
	queue []interface{}
	size  int
	mu    sync.Mutex
}

func newSafeQueue(n int) *safeQueue {
	return &safeQueue{size: n}
}

func (q *safeQueue) isEmpty() bool {
	return len(q.queue) == 0
}

func (q *safeQueue) isFull() bool {
	return len(q.queue) == q.size
}

func (q *safeQueue) front() (interface{}, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.isEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.queue[0], nil
}

func (q *safeQueue) enqueue(val interface{}) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.isFull() {
		return errors.New("queue is full")
	}
	q.queue = append(q.queue, val)
	return nil
}

func (q *safeQueue) dequeue() error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.isEmpty() {
		return errors.New("queue is empty")
	}
	q.queue = q.queue[1:]
	return nil
}

func main() {
	queue := NewSafeQueue(10)

	// Enqueue items.
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)

	// Dequeue items.

	for i := 0; i < len(queue.queue); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			item, err := queue.front()
			if err != nil {
				fmt.Println("error")
			}
			fmt.Println(item)
			queue.dequeue()
		}()

	}
	wg.Wait()
	fmt.Println("DONE")

}
