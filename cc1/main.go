package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 3}

	fmt.Println(GetSymmetric(a, b))
}

func GetSymmetric(a, b []int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(a); i++ {

		m[a[i]] += 1

	}
	for i := 0; i < len(b); i++ {

		m[b[i]] += 1

	}
	res := []int{}
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
