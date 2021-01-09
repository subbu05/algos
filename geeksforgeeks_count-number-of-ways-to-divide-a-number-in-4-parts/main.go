package main

import "fmt"

func main() {
	fmt.Println(getWaysWorst(6))
}

// )(n power 4 ) complexity
func getWaysWorst(n int) int {
	count := 0
	for i := 1; i < n; i++ {
		for j := i; j < n; j++ {
			for k := j; k < n; k++ {
				for l := k; l < n; l++ {
					if i+j+k+l == n {
						count++
					}
				}
			}
		}
	}

	return count
}
