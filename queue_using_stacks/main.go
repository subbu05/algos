package main

import "fmt"

type stack struct {
	elm  []int
	size int
	top  int
}

var sq *stack

func enqueue(s *stack, i int) {
	if (s.top) < s.size {
		s.top += 1
		s.elm[s.top] = i
	}
}

func dequeue(s *stack) {
	if len(s.elm) > 0 {
		fmt.Println(s.elm[len(s.elm)-1])
		helper(s, sq)
	}
}

func helper(from, to *stack) {
	if len(from.elm) > 0 {
		i := 0
		for j := len(from.elm) - 1; j > 0; j-- {
			to.elm[i] = from.elm[j]
			i++
		}
	}
	for len(to.elm) > 0 {
		i := 0
		for j := len(to.elm) - 1; j >= 0; j-- {
			from.elm[i] = to.elm[j]
			i++
		}
	}
}

func main() {
	sq = &stack{
		elm:  make([]int, 5),
		size: 0,
		top:  -1,
	}

	q := &stack{
		elm:  make([]int, 5),
		size: 5,
		top:  -1,
	}

	enqueue(q, 5)
	fmt.Println(q.elm)
}
