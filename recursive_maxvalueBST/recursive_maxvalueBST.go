package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func main() {
	root := buildtree()
	fmt.Println(MaxBST(root))
}

func MaxBST(root *Node) int {
	if root != nil {
		if root.Left == nil {
			return root.Data
		}
		return MaxBST(root.Right)
	}
	return -1
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
