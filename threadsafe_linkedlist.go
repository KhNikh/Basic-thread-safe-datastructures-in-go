package main

import (
	"fmt"
	"sync"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	mu   sync.Mutex // Use RWMutex for read/write locking
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) SafeInsert(value int) {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	newNode := &Node{Value: value}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (ll *LinkedList) Print() {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	current := ll.Head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	ll := NewLinkedList()

	// Concurrently insert values into the linked list
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			ll.SafeInsert(val)
		}(i)
	}
	wg.Wait()

	ll.Print()
}
