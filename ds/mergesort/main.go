package main

import "fmt"

func main() {
	a := []int{5, 4, 3, 2, 1}
	mergseSort(a, 0, len(a))
	fmt.Println(a)
}

func mergseSort(a []int, s, e int) {
	if s >= e {
		return
	}
	m := (s + e) / 2
	mergseSort(a, s, m)
	mergseSort(a, m+1, e)
	merge(a, s, m, e)
}

func merge(a []int, s, m, e int) {
	b := make([]int, m-s)
	c := make([]int, e-m)
	for i := s; i < m; i++ {
		b[i] = a[i]
	}
	for i := m; i < e; i++ {
		c[i] = c[i]
	}

}
