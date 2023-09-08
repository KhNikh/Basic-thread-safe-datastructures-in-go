package main

import (
	"errors"
	"fmt"
)

type stack struct {
	stack []interface{}
	size  int
}

func (s *stack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *stack) isFull() bool {
	return len(s.stack) == s.size
}

func (s *stack) push(val interface{}) error {
	if s.isFull() {
		return errors.New("Stack is full")
	}
	s.stack = append(s.stack, val)
	return nil
}

func (s *stack) pop() error {
	if s.isEmpty() {
		return errors.New("Stack is empty")
	}
	s.stack = s.stack[:len(s.stack)-1]
	return nil
}

func (s *stack) Peek() (interface{}, error) {
	if s.isEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.stack[len(s.stack)-1], nil
}

func main() {
	s := stack{
		size: 10,
	}
	err := s.pop()
	if err != nil {
		fmt.Println("Error:", err)
	}
	s.push(12)
	top, err := s.Peek()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The top element is ", top)
	}
}
