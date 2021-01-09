package main

import "fmt"

type node struct {
	Data  int
	Left  *node
	Right *node
}

func main() {
	root := &node{
		Data:  10,
		Left:  nil,
		Right: nil,
	}

	root.Left = &node{
		Data:  5,
		Left:  nil,
		Right: nil,
	}

	root.Right = &node{
		Data:  15,
		Left:  nil,
		Right: nil,
	}

	postorderIterative(root)
}

func postorderIterative(root *node) {
	if root == nil {
		return
	}
	s := []*node{}
	out := []int{}

	s = append(s, root)
	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[:len(s)-1]
		out = append(out, cur.Data)

		if cur.Left != nil {
			s = append(s, cur.Left)
		}
		if cur.Right != nil {
			s = append(s, cur.Right)
		}

	}

	for i := len(out) - 1; i >= 0; i-- {
		fmt.Println(out[i])
	}
}
