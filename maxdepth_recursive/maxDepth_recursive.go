package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func main() {
	root := buildtree()
	fmt.Println(maxDepth(root))
}

func buildtree() *Node {
	root := &Node{
		Data:  5,
		Left:  nil,
		Right: nil,
	}

	root.Left = &Node{
		Data:  3,
		Left:  nil,
		Right: nil,
	}

	root.Right = &Node{
		Data:  7,
		Left:  nil,
		Right: nil,
	}

	root.Left.Left = &Node{
		Data:  2,
		Left:  nil,
		Right: nil,
	}

	return root
}

func maxDepth(root *Node) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	l := 1 + maxDepth(root.Left)
	r := 1 + maxDepth(root.Right)
	if l < r {
		return r
	}
	return l
}
