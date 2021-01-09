package main

import "fmt"

func main() {

	fmt.Println(helper())
}

func helper() int {
	a := [][]int{{123, 6, 1}, {124, 5, 1}}
	m := make(map[int]int, 0)
	max := 0

	for i := 0; i < len(a); i++ {
		if a[i][2] == 0 {
			m[a[i][0]] -= a[i][1]
		} else {
			m[a[i][0]] += a[i][1]
		}
		//fmt.Println(a[i][0], m[a[i][0]])
		if m[a[i][0]] > max {
			max = m[a[i][0]]
		}
		//fmt.Println(max)
	}

	for k, v := range m {
		if v == max {
			return k
		}
	}

	return 0
}
