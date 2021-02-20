package main

type node struct {
	val  int
	next *node
}

func main() {

}

func reverseListIterative(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	cur := head
	var prev *node = nil
	for cur != nil {
		temp := cur.next
		cur.next = prev
		prev = cur
		cur = temp
	}
	return prev
}

func reverseListRecursive(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	p := reverseListRecursive(head.next)
	head.next.next = head
	head.next = nil
	return p
}
