package main

import "fmt"

var A []int

func main() {
	A = make([]int, 2)
	RecBin(2)
}

func RecBin(n int) {
	if n <= 1 {
		fmt.Println(A)
		return
	}

	A[n-1] = 0
	RecBin(n - 1)
	A[n-1] = 1
	RecBin(n - 1)
}
