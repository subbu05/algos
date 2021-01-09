package main

import "fmt"

type node struct {
	Val  int
	Next *node
}

func main() {
	var head *node = nil
	head = insertFront(head, 1)
	head = insertFront(head, 2)
	head = insertFront(head, 3)
	view(head)
	fmt.Println(nthNode(head, 3))
	fmt.Println(recursiventhNode(head, 3))
}

func nthNode(head *node, pos int) *node {
	if pos <= 0 {
		return nil
	}
	fast := head
	var slow *node = nil
	for i := 1; i < pos; i++ {
		if fast != nil {
			fast = fast.Next
		}
	}

	for fast != nil {
		if slow == nil {
			slow = head
		} else {
			slow = slow.Next
		}
		fast = fast.Next
	}
	return slow
}

var counter = 0

func recursiventhNode(head *node, nth int) *node {
	if head != nil {
		recursiventhNode(head.Next, nth)
		counter++
		if counter == nth {
			return head
		}
	}
	return nil
}

func insertFront(head *node, v int) *node {

	if head != nil {
		return &node{
			Val:  v,
			Next: head,
		}
	}

	return &node{
		Val: v,
	}
}

func view(head *node) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
