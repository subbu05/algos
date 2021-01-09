package main

import "fmt"

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func main() {
	root := &Tree{
		Val: 0,
		Left: &Tree{
			Val: 5,
			Left: &Tree{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &Tree{
			Val:   10,
			Left:  nil,
			Right: nil,
		},
	}

	Inorder(root)
}

func Inorder(root *Tree) {
	if root == nil {
		return
	}
	s := []*Tree{}
	cur := root

	for cur != nil || len(s) > 0 {
		for cur != nil {
			s = append(s, cur)
			cur = cur.Left
		}

		cur = s[len(s)-1]
		s = s[:len(s)-1]

		fmt.Print(cur.Val)

		cur = cur.Right
	}
}
