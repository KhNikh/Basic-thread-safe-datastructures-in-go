package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var wr sync.WaitGroup

type safeStack struct {
	stack []interface{}
	mu    sync.Mutex
}

func newSafeStack() *safeStack {
	return &safeStack{}
}

func (s *safeStack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *safeStack) peek() (interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.isEmpty() {
		return nil, errors.New("Stack is empty")
	}
	return s.stack[len(s.stack)-1], nil
}

func (s *safeStack) pop() (interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.isEmpty() {
		return nil, errors.New("Stack is empty")
	}
	item := s.stack[len(s.stack)-1]
	if len(s.stack) > 1 {
		s.stack = s.stack[:(len(s.stack) - 2)]
	} else {
		s.stack = s.stack[:]
	}
	return item, nil
}

func (s *safeStack) push(val int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stack = append(s.stack, val)
}

func main() {
	stack := newSafeStack()

	// Enqueue items.
	stack.push(1)
	stack.push(2)
	stack.push(3)

	// Dequeue items.

	for i := 0; i < len(stack.stack); i++ {
		wr.Add(1)
		go func() {
			defer wr.Done()
			item, err := stack.peek()
			if err != nil {
				fmt.Println("error")
			}
			fmt.Println(item)
			item, err = stack.pop()
			if err != nil {
				fmt.Println("error")
			}
		}()
		time.Sleep(1 * time.Second)

	}
	wr.Wait()
	fmt.Println("DONE")

}
