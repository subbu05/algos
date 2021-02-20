package main

import "fmt"

type node struct {
	val  int
	Next *node
}

func main() {
	head := &node{
		val: 1,
		Next: &node{
			val: 2,
			Next: &node{
				val: 3,
				Next: &node{
					val: 4,
					Next: &node{
						val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	head = swapNodes(head, 2, 3)
	viewList(head)
}

func viewList(head *node) {
	for head != nil {
		fmt.Println(head.val)
		head = head.Next
	}
}

func swapNodes(head *node, m, n int) *node {
	if head == nil || head.Next == nil || m == n {
		return head
	}

	slow := head
	var prev *node

	for m > 1 && slow != nil {
		prev = slow
		slow = slow.Next
		m--
		n--
	}
	//fmt.Println(slow.val)

	cur := slow
	con := prev

	for n > 0 && cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
		n--
	}

	if con != nil {
		con.Next = prev
	} else {
		head = prev
	}
	slow.Next = cur
	return head
}
