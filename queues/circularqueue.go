package main

import "fmt"

const max = 5

var head = -1
var tail = -1
var q = [max]int{}

func main() {

	fmt.Println(q)
	add(5)
	remove()
	fmt.Println(head)
	fmt.Println(tail)
}

func add(d int) {
	if !isFull() {
		tail = (tail + 1) % max
		q[tail] = d

		if head == -1 {
			head = tail
		}
	}
}

func remove() {
	if !isEmpty() {
		fmt.Println(q[head])
		if head == tail {
			head = -1
		} else {
			head = (head + 1) % max
		}
	}
}

func isEmpty() bool {
	return head == -1
}

func isFull() bool {
	temp := (tail + 1) % max
	return temp == head
}
