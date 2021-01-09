package main

import "fmt"

func main() {
	fmt.Print(isArraySorted([]int{1, 23, 3, 4, 5}, 4))
}

func isArraySorted(a []int, n int) bool {
	if len(a) <= 1 || n == 0 {
		return true
	}

	if a[n] > a[n-1] {
		return isArraySorted(a, n-1)
	}

	return false
}
