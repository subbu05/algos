package main

import "fmt"

func main() {
	fmt.Println(sum2(1050))
}

func sum(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if i%5 == 0 {
			sum += i
		} else if i%3 == 0 {
			sum += i
		}
	}
	return sum
}

func sum2(n int) int {
	t := 999 / n
	return n * ((t * (t + 1)) / 2)
}
