package main

import "fmt"

type node struct {
	val  int
	next *node
}

func main() {
	var head *node = nil
	head = insertIntoSortedList(head, 3)
	view(head)
	head = insertIntoSortedList(head, 5)
	head = insertIntoSortedList(head, 4)
	head = insertIntoSortedList(head, 3)
	view(head)
}

func view(head *node) {
	cur := head
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}

func insertIntoSortedList(head *node, i int) *node {
	temp := &node{
		val:  i,
		next: nil,
	}
	if head == nil {
		return temp
	} else {
		var prev *node = nil
		cur := head
		for cur != nil && cur.val < i {
			prev = cur
			cur = cur.next
		}
		if prev == nil {
			temp.next = head
			return temp
		} else {
			prev.next = temp
			temp.next = cur
		}
	}
	return head
}
