package main

import "fmt"

type node struct {
	val  int
	next *node
}

type stack struct {
	head *node
	top  int
	max  int
}

func constructor(size int) *stack {
	s := stack{
		head: nil,
		top:  -1,
		max:  size,
	}
	return &s
}

func main() {
	s := constructor(5)
	push(s, 1)
	push(s, 2)
	view(s)
}

func peek(s *stack) int {
	if s.top == -1 {
		return -1
	}
	return s.head.val
}

func push(s *stack, i int) {
	if !isFull(s) {
		temp := &node{
			val:  i,
			next: s.head,
		}
		s.head = temp
	}
}

func pop(s *stack) {
	if !isEmpty(s) {
		s.head = s.head.next
	}
}

func isEmpty(s *stack) bool {
	return s.top == -1
}

func isFull(s *stack) bool {
	return s.top == s.max
}

func view(s *stack) {
	cur := s.head
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}
