package main

import "fmt"

type tree struct {
	val int
	left *tree
	right *tree
}


func main() {
	head := &tree{
		val:   50,
		left:  &tree{
			val:   7,
			left:  &tree{
				val:   2,
				left:  nil,
				right: nil,
			},
			right: &tree{
				val:   1,
				left:  nil,
				right: &tree{
					val:   8,
					left:  nil,
					right: &tree{
						val:   0,
						left:  nil,
						right: nil,
					},
				},
			},
		},
		right: nil,
	}

	fmt.Print(findMaxIterative(head))
}

func findMaxIterative(root *tree) int {
	max := -1
	if root == nil {
		return max
	}
	s := []*tree{root}

	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[0:len(s)-1]
		if cur.val > max {
			max = cur.val
		}
		if cur.right != nil {
			s = append(s,cur.right)
		}
		if cur.left != nil {
			s = append(s, cur.left)
		}
	}
	return max
}

