package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 5, 5, 5}
	max := 0
	n := len(a)
	for i := 0; i < len(a); i++ {
		a[a[i]%n] += n
	}
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		if a[i]/n > max {
			max = a[i] / n
		}
	}
	fmt.Println(max)
}
