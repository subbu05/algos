package main

import (
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/push"
)

type singleLinkedList struct {
	val  int
	next *singleLinkedList
}

var EMPTY_STACK_ERROR error = errors.New("Empty stack")

func insertFrontSingleLinkedList(head *singleLinkedList, v int) *singleLinkedList {
	temp := &singleLinkedList{
		val:  v,
		next: nil,
	}
	if head != nil {
		return temp
	}
	temp.next = head
	return temp
}

func deleteFrontSingleLinkedList(head *singleLinkedList) (*singleLinkedList, int) {
	if head == nil {
		return nil
	}
	return head.next
}

func main() {
	sllAsStack()
}

type stack struct {
	head *singleLinkedList
	top  int
}

func peekStack(s *stack) (int, error) {
	if s.top >= 0 {
		return s.head.val, nil
	}
	return 0, EMPTY_STACK_ERROR
}

func (s *stack) push(v int) {
	s.head = insertFrontSingleLinkedList(s.head, v)
}

/*func (s *stack) pop() (int,error) {
    deleteFrontSingleLinkedList(s.head)

}*/

func sllAsStack() {
	s := &stack{
		head: nil,
		top:  -1,
	}

}
