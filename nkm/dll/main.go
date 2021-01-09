package main

import "fmt"

type node struct {
	Val  int
	Prev *node
	Next *node
}

func main() {
	var head *node
	head = insertFront(head, 5)
	head = insertBack(head, 55)
	head = insertBack(head, 555)

	head = deleteFront(head)
	head = deleteBack(head)
	head = insertBack(head, 7)
	head = insertFront(head, 3)
	view(head)
}

func insertFront(head *node, v int) *node {
	temp := &node{
		Val:  v,
		Prev: nil,
		Next: nil,
	}
	if head == nil {
		head = temp
		return head
	}
	temp.Next = head
	head.Prev = temp
	//head.Next = nil
	head = temp
	return head
}

func deleteFront(head *node) *node {
	if head == nil || (head.Prev == nil && head.Next == nil) {
		return nil
	}
	head = head.Next
	head.Prev = nil
	return head
}

func insertBack(head *node, v int) *node {
	if head == nil {
		return &node{
			Val:  v,
			Prev: nil,
			Next: nil,
		}
	}
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &node{
		Val:  v,
		Prev: cur,
		Next: nil,
	}
	return head
}

func deleteBack(head *node) *node {
	if head == nil || (head.Prev == nil && head.Next == nil) {
		return nil
	}
	cur := head
	for cur.Next != nil && cur.Next.Next != nil {
		cur = cur.Next
	}
	cur.Next = nil
	return head
}

func view(head *node) {
	if head != nil {
		cur := head
		for cur != nil {
			fmt.Println(cur.Val)
			cur = cur.Next
		}
	}
}
