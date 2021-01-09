package main

import (
	"fmt"
)

type node struct {
	Val  int
	Next *node
}

func main() {
	var head *node = nil
	head = insertFront(head, 5)
	view(head)
}

func insertFront(head *node, v int) *node {
	if head == nil {
		head = &node{
			Val:  v,
			Next: nil,
		}
		head.Next = head
		return head
	}
	temp := &node{
		Val:  0,
		Next: head,
	}
	cur := head.Next
	for cur != head {
		cur = cur.Next
	}
	cur.Next = temp
	return head
}

func view(head *node) {
	if head != nil {
		cur := head.Next
		for cur != head {
			fmt.Println(cur.Val)
			cur = cur.Next
		}
	}
}
