package main

import "fmt"

var q [5]int
var head = -1
var tail = -1
var N = 5

func main() {
	q = [5]int{}
	enqueue(5)
	enqueue(4)
	dequeue()
	enqueue(3)
	dequeue()
}

func enqueue(elm int) {
	if !isFull() {
		tail = (tail + 1) % N
		q[tail] = elm
	}

	if head == -1 {
		head = tail
	}
}

func dequeue() {
	if !isEmpty() {
		fmt.Println(q[head])
		if head == tail {
			head = -1
		} else {
			head = (head + 1) % N
		}

	}
}

func isFull() bool {
	i := (tail + 1) % N
	return i == head
}

func isEmpty() bool {
	return head == -1
}
