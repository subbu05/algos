package main

import "fmt"

func main() {
	a := []int{5, 1, 3, 2, 4}
	shellsort(a)
	fmt.Println(a)
}

func shellsort(a []int) {

	inc := len(a) / 2

	for inc > 0 {
		for i := 0; i < inc; i++ {
			inSort(a, i, inc)
		}
		inc /= 2
	}

}

func inSort(a []int, start, inc int) {
	j := 0
	for i := start; i < len(a); i += inc {
		temp := a[i]
		for j = i - inc; j >= 0 && a[j] > temp; j = j - inc {
			a[j+inc] = a[j]
		}
		a[j+inc] = temp
	}
}
