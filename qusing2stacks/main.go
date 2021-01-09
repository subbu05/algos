package main

import "fmt"

type Queue struct {
	elms  []int
	front int
	rear  int
	max   int
}

var Stack1 []int
var Stack2 []int

func construct() Queue {
	return Queue{
		elms:  nil,
		front: -1,
		rear:  -1,
		max:   5,
	}
}

func Enqueue(q Queue, d int) {
	if len(Stack1) <= q.max {
		Stack1 = append(Stack1, d)
	}
}

func Dequeue() {
	for i := len(Stack1) - 1; i > 0; i-- {
		Stack2 = append(Stack2, Stack1[i])
	}
	fmt.Println(Stack1[0])
	for i := len(Stack1) - 1; i > 0; i-- {
		Stack1 = append(Stack1, Stack2[i])
	}

}

func main() {
	q := construct()
	Stack1 = make([]int, q.max)
	Stack2 = make([]int, q.max)

}

func Peek(queue Queue) {
	if len(Stack1) > 0 {
		fmt.Println(Stack1[0])
	}
}
