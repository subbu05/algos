package main

import "fmt"

type node struct {
	Val  int
	Next *node
}

func main() {
	var head *node = nil
	head = insertFront(head, 1)
	head = insertFront(head, 3)
	head = insertBack(head, 5)
	head = deleteFront(head)
	head = deleteBack(head)
	//head = deleteBack(head)
	view(head)
}

func view(head *node) {
	//fmt.Println(head)
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func insertFront(head *node, v int) *node {
	temp := node{
		Val:  v,
		Next: head,
	}
	head = &temp
	return head
}

func deleteFront(head *node) *node {
	if head != nil {
		head = head.Next
	}
	return head
}

func insertBack(head *node, v int) *node {
	temp := &node{
		Val:  v,
		Next: nil,
	}

	if head == nil {
		return temp
	}

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = temp
	return head
}

func deleteBack(head *node) *node {
	if head == nil || head.Next == nil {
		return nil
	}
	cur := head
	var prev *node
	for cur.Next != nil {
		prev = cur
		cur = cur.Next
	}
	prev.Next = nil
	return head
}

func insertPos(head *node, v, pos int) *node {
	if pos == 1 {
		temp := &node{
			Val:  v,
			Next: head,
		}
		return temp
	}
	c := 0
	cur := head

	for c < pos && cur != nil {
		c++
		cur = cur.Next
	}
	if pos == c && cur != nil {
		cur.Next = &node{
			Val: v,
		}
		return head
	}
	return head
}

func deletePos(head *node, pos int) *node {
	if head == nil {
		return nil
	}
	c := 0
	cur := head
	prev := head
	for c != pos && cur != nil {
		c++
		prev = cur
		cur = cur.Next
	}

}
