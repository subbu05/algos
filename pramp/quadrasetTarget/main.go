package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getQuads([]int{4, 2, 3, 1}, 10))
}

func getQuads(a []int, s int) []int {
	n := len(a)
	r := 0
	l := 0
	h := 0
	if n < 4 {
		return []int{}
	}
	sort.Ints(a)
	for i := 0; i <= n-4; i++ {
		for j := i + 1; j <= n-3; j++ {
			l = j + 1
			h = n - 1
			r = s - (a[i] + a[j])
			for l < h {
				if a[l]+a[h] < r {
					l++
				} else if a[l]+a[h] > r {
					h--
				} else {
					return []int{a[i], a[j], a[l], a[h]}
				}
			}
		}
	}
	return []int{}
}
