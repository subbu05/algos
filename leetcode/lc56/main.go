package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{2, 3}, {1, 6}}))
}

func merge(intervals [][]int) [][]int {

	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][0] < intervals[i][1] })

	res := [][]int{}

	for _, i := range intervals {
		fmt.Println(res)
		if len(res) == 0 || res[len(res)-1][1] < i[0] {
			res = append(res, i)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], i[1])
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
