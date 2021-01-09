package main

import "fmt"

func main() {
	a := []int{5, 1, 2, 3, 4}
	mergesort(a, 0, len(a)-1)
	fmt.Println(a)
}

func mergesort(a []int, low, high int) {

	if low < high {
		mid := (low + high) / 2
		mergesort(a, low, mid)
		mergesort(a, mid+1, high)
		merge(a, low, mid, high)
	}
}

func merge(a []int, low, mid, high int) {
	fmt.Println(a)
	fmt.Println(low)
	fmt.Println(mid)
	fmt.Println(high)
	return
	b := a[low : mid+1]
	c := a[mid+1 : high+1]
	k := 0
	i := 0
	j := 0

	for i < len(b) && j < len(c) {
		if b[i] < c[j] {
			a[k] = b[i]
			k++
			i++
		} else {
			a[k] = c[j]
			k++
			j++
		}
	}

	for i < len(b) {
		a[k] = b[i]
		k++
		i++
	}

	for j < len(c) {
		a[k] = b[j]
		k++
		j++
	}

}
