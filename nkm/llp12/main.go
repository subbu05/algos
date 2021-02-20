package main

import "fmt"

type node struct {
	val  int
	next *node
}

func main() {
	head := &node{
		val: 0,
		next: &node{
			val: 2,
			next: &node{
				val:  0,
				next: nil,
			},
		},
	}
	head.next.next = head
	fmt.Println(helper(head))
}

func helper(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	slow := head
	fast := head
	loopExists := false
	for fast != nil || fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			loopExists = true
			break
		}
	}

	if loopExists {
		slow = head
		for slow != fast {
			slow = slow.next
			fast = fast.next
		}
		return fast
	}
	return nil
}
