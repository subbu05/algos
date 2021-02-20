package main

type node struct {
	val  int
	next *node
}

func main() {
}

func reverseInPairsRecursive(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	temp := head.next
	head.next = temp.next
	temp.next = head
	head = temp
	head.next.next = reverseInPairsRecursive(head.next.next)
	return head
}

func reverseInPairsIterative(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	var t1 *node
	var t2 *node
	for head != nil && head.next != nil {
		if t1 != nil {
			t1.next.next = head.next
		}
		t1 = head.next
		head.next = head.next.next
		t1.next = head
		if t2 == nil {
			t2 = t1
		}
		head = head.next
	}
	return t2
}
