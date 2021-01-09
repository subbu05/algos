package main

import "fmt"

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func main() {
	root := &Tree{
		Val: 7,
		Left: &Tree{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &Tree{
			Val: 3,
			Left: &Tree{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &Tree{
				Val:   19,
				Left:  nil,
				Right: nil,
			},
		},
	}

	cloned := root

	fmt.Println(getCloneTarget(root, cloned, root.Right))
}

func getCloneTarget(o, c, t *Tree) *Tree {
	if o == nil || t == nil {
		return o
	}
	s1 := []*Tree{o}
	s2 := []*Tree{c}

	for len(s1) > 0 && len(s2) > 0 {
		cur1 := s1[len(s1)-1]
		cur2 := s2[len(s2)-1]
		s1 = s1[0 : len(s1)-1]
		s2 = s2[0 : len(s2)-1]

		if cur1 == t {
			return cur2
		}

		if cur1.Left != nil && cur2 != nil {
			s1 = append(s1, cur1.Left)
			s2 = append(s2, cur2.Left)
		}

		if cur1.Right != nil && cur2.Right != nil {
			s1 = append(s1, cur1.Right)
			s2 = append(s2, cur2.Right)
		}

	}
	return nil
}
