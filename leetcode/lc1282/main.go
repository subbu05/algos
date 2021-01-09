package main

import "fmt"

func main() {
	fmt.Print(groupThePeople([]int{3, 3, 3, 1, 3}))
}

func groupThePeople(groupSizes []int) [][]int {
	res := [][]int{}
	tmp := make(map[int][]int)

	for i, v := range groupSizes {
		tmp[v] = append(tmp[v], i)
		if len(tmp[v]) == v {
			res = append(res, tmp[v])
			tmp[v] = []int{}
		}
	}
	return res
}
