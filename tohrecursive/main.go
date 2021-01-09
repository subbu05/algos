package main

import "fmt"

func main() {
	toh("From", "To", "Aux", 5)
}

func toh(F, T, A string, n int) {
	if n == 1 {
		fmt.Printf("Move Disc  from %s to % s\n", F, T)
		return
	}

	fmt.Printf("Move disc  %s to %s\n", F, A)
	toh(F, A, T, n-1)

	fmt.Printf("Move disc  %s to %s\n", A, T)
	toh(A, T, F, n-1)
}
