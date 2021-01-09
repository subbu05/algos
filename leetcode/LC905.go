package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 5, 6, 8}))
}

func sortArrayByParity(A []int) []int {
	i := 0
	j := len(A) - 1

	for i < j {
		fmt.Printf("%d %d\n", i, j)
		if A[i]%2 > A[j]%2 {
			fmt.Println(A)
			fmt.Println("in CMP")
			A[i], A[j] = A[j], A[i]

		}

		fmt.Println(A)

		if A[i]%2 == 0 {
			i++

		}
		fmt.Println(i)

		if A[j]%2 != 0 {
			j--
		}
		fmt.Println(j)
		os.Exit(0)
	}
	return A
}
