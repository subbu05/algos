package main

// TC/SC O(max(m,n))
type node struct {
	Val  int
	Next *node
}

func addTwoNumbers(l1 *node, l2 *node) *node {
	l3 := &node{}
	cur := l3
	c := 0
	sum := 0
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + c
		temp := &node{}
		temp.Val = sum % 10
		c = sum / 10
		l3.Next = temp
		l3 = l3.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum = l1.Val + c
		temp := &node{}
		temp.Val = sum % 10
		c = sum / 10
		l3.Next = temp
		l3 = l3.Next
		l1 = l1.Next

	}
	for l2 != nil {
		sum = l2.Val + c
		temp := &node{}
		temp.Val = sum % 10
		c = sum / 10
		l3.Next = temp
		l3 = l3.Next
		l2 = l2.Next

	}

	if c != 0 {
		temp := &node{}
		temp.Val = c
		l3.Next = temp
	}
	return cur.Next
}

func main() {
	l1 := &node{
		Val:  5,
		Next: nil,
	}
	l2 := &node{
		Val:  9,
		Next: nil,
	}
	addTwoNumbers(l1, l2)
}
