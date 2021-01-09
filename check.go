package main

import "fmt"

func main() {

	q := []int{1, 2, 3}

	for len(q) > 0 {
		cur := q[0]
		fmt.Println(cur)
		q = q[1:]
	}
}
