package main

import "fmt"

type stack []int

var s stack
var ms stack

func main() {

	s = make(stack, 0)
	ms = make(stack, 0)

	push(9)
	push(2)
	push(3)
	push(1)
	pop()
	fmt.Println(getMin())
}

func push(i int) {
	if !isFull(s) {
		s = append(s, i)
		if !isEmpty(ms) {
			if ms[len(ms)-1] > i {
				ms = append(ms, i)
			} else {
				ms = append(ms, ms[len(ms)-1])
			}
		} else {
			ms = append(ms, i)
		}
	}
}

func pop() {
	if !isEmpty(s) {
		s = s[:len(s)-1]
		ms = ms[:len(ms)-1]
	}
}

func getMin() int {
	if len(ms) < 1 {
		return -1
	}

	return ms[len(ms)-1]
}

func isEmpty(s stack) bool {
	return len(s) < 1
}

func isFull(s stack) bool {
	return len(s) > 5
}
