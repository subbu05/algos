package main

import "fmt"

func main() {
	a := []int{5, 4, 3, 2, 1}
	bubblesort(a)
	fmt.Println(a)
}

func bubblesort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	//fmt.Println(a)
}
