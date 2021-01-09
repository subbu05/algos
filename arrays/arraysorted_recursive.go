package main

import "fmt"

func main() {
	fmt.Println(isArraySorted([]int{5, 4, 55, 2, 1}, 0))
}

func isArraySorted(a []int, i int) bool {
	if i == len(a)-1 {
		return true
	}
	if a[i] > a[i+1] == false {
		return false
	}
	return isArraySorted(a, i+1)
}
