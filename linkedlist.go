package main

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}
type linkedlist struct {
	head *node
}

func newLinkedList() *linkedlist {
	return &linkedlist{}
}
func (ll *linkedlist) lengthOfLL() int {
	count := 0
	for it := ll.head; it != nil; it = it.next {
		count++
	}

	return count
}

func (ll *linkedlist) printList() {
	if ll.head == nil {
		fmt.Println("List is empty")
	}
	for it := ll.head; it != nil; it = it.next {
		fmt.Println(it.val)
	}
}

func (ll *linkedlist) insertNode(data int, idx int) {
	if idx < 0 || idx > (ll.lengthOfLL()+1) {
		fmt.Println("Index does not exist")
	} else {
		it := ll.head
		for i := 0; i < idx-2; i++ {
			it = it.next
		}
		if it == nil {
			newNode := &node{val: data, next: nil}
			ll.head = newNode

		} else {
			if it.next == nil {
				newNode := &node{val: data, next: nil}
				it.next = newNode
			} else {
				newNode := &node{val: data, next: it.next}
				it.next = newNode
			}
		}

	}
}

func (ll *linkedlist) deleteNode(idx int) {

	if idx < 0 || ll.lengthOfLL() < idx {
		fmt.Println("Index does not exist")
	} else {
		it := ll.head
		for i := 0; i < idx-2; i++ {
			it = it.next
		}
		temp := it.next
		it.next = temp.next
		temp.next = nil
	}
}

func main() {

	ll := newLinkedList()
	ll.insertNode(2, 1)
	ll.insertNode(6, 2)
	ll.insertNode(3, 2)
	ll.insertNode(4, 2)
	ll.printList()
	fmt.Println("===============================")
	ll.deleteNode(2)
	ll.printList()

}
