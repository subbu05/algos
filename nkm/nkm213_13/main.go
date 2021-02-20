package main

import "fmt"

func main() {
	fmt.Println(getMid([]int{1, 2, 3, 4, 5, 8, 6, 9}))
}

func getMid(a []int) int {
	if len(a) == 0 {
		return -1
	}
	sum := 0
	lsum := 0

	for _, v := range a {
		sum += v
	}

	for i := 0; i < len(a); i++ {
		sum -= a[i]
		if lsum == sum {
			return a[i]
		}
		lsum += a[i]
	}

	return -1
}
