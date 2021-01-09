package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func main() {
	head := &Node{
		Val: 1,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val: 3,
				Next: &Node{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	ViewList(head)
	RemoveBack(head)
	RemoveBack(head)
	RemoveBack(head)
	RemoveBack(head)
	ViewList(head)
}

func SizeInterative(head *Node) int {
	if head == nil {
		return 0
	}

	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}

	return count
}

func SizeRecursive(head *Node) int {
	if head == nil {
		return 0
	}

	return 1 + SizeRecursive(head.Next)
}

func InsertFront(head *Node, v int) {
	temp := &Node{
		Val: v,
	}
	if head == nil {
		head = temp
		return
	}

	temp.Next = head
	head = temp
}

func RemoveFront(head *Node) {
	if head == nil {
		return
	}

	head = head.Next
}

func InsertBack(head *Node, v int) {
	temp := &Node{
		Val: v,
	}
	if head == nil {
		head = temp
		return
	}

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = temp
}

func RemoveBack(head *Node) {
	if head == nil || head.Next == nil {
		head = nil
		return
	}
	cur := head
	var prev *Node
	for cur.Next != nil {
		prev = cur
		cur = cur.Next
	}
	prev.Next = nil
}

func ViewList(head *Node) {
	if head == nil {
		return
	}
	cur := head
	for cur != nil {
		fmt.Printf("%d==>", cur.Val)
		cur = cur.Next
	}
	fmt.Println("NIL")
}
