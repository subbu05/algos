package main

type node struct {
	val  int
	next *node
}

func main() {

}

func reverseList(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	var prev, temp *node
	cur := head

	for cur != nil {
		temp = cur.next
		cur.next = prev
		prev = cur
		cur = temp
	}
	return prev
}
