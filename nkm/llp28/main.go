package main

import "fmt"

type node struct {
	val  int
	next *node
}

func main() {
	head := &node{
		val: 1,
		next: &node{
			val: 2,
			next: &node{
				val: 3,
				next: &node{
					val:  4,
					next: nil,
				},
			},
		},
	}
	temp := getMiddle(head)
	if temp != nil {
		fmt.Println(temp.val)
	}
}

func getMiddle(head *node) *node {
	if head == nil || head.next == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
