package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func main() {

}

func PostOrder(root *Tree) {
	if root == nil {
		return
	}
	s1 := []*Tree{root}
	s2 := []*Tree{root}

	for len(s1) > 0 {
		cur := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]
		s2 = append(s2, cur)

		if cur.Left != nil {
			s1 = append(s1, cur.Left)
		}
		if cur.Right != nil {
			s1 = append(s1, cur.Right)
		}
	}

	for i := len(s2) - 1; i >= 0; i-- {
		fmt.Print(s2[i])
	}

	rand.Intn(5)

}
