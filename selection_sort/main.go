package main

import "fmt"

func main() {
	a := []int{5, 4, 3, 2, 1}

	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}

	fmt.Println(a)
}
