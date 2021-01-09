package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}
type Stack struct {
	Top  int
	Max  int
	Elms *Node
}

func main() {

}

func Push(stack Stack, d int) {
	if !isFull(stack) {
		temp := &Node{
			Val:  d,
			Next: nil,
		}
		if stack.Elms == nil {

			stack.Elms = temp
		} else {
			temp.Next = stack.Elms
			stack.Elms = temp
		}
		stack.Top++
	}
}

func Pop(stack Stack) {
	if !isEmpty(stack) {
		stack.Elms = stack.Elms.Next
		stack.Top--
	}
}

func Peek(stack Stack) {
	if stack.Top > -1 {
		fmt.Println(stack.Elms.Val)
	}
}

func isEmpty(stack Stack) bool {
	return stack.Top == -1
}

func isFull(stack Stack) bool {
	return stack.Top >= stack.Max
}
