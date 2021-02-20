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

func helper(head *node) int {
	if head == nil || head.next == nil {
		return 0
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

	c := 0
	if loopExists {
		slow = slow.next
		c = 1
		for slow != fast {
			slow = slow.next
			c++
		}
	}
	return c
}
