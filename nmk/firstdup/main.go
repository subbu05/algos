package main

import "fmt"

func main() {
	a := []int{3, 2, 1, 2, 2, 3}
	fmt.Println(UsingMap(a))
}

func BruteForceMethod(a []int) bool {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				fmt.Println(a[i])
				return true
			}
		}
	}
	return false
}

func UsingMap(a []int) bool {
	m := make(map[int]int, 0)
	for i := 0; i < len(a); i++ {
		if v, ok := m[a[i]]; ok {
			m[a[i]] = -v
		} else {
			m[a[i]] = i + 1
		}
	}
	min := 0
	for k, v := range m {
		if v >= 0 {
			delete(m, k)
		} else {
			if v < m[min] {
				min = k
			}
		}

	}
	fmt.Println(min)
	return true
}
