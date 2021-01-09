package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayRankTransform([]int{-40, 10, 20, 30}))
}
func arrayRankTransform(arr []int) []int {
	copyArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		copyArr[i] = arr[i]
	}
	sort.Ints(copyArr)
	m := make(map[int]int, 0)
	rank := 1
	for i := 0; i < len(copyArr); i++ {
		if _, ok := m[copyArr[i]]; ok == false {
			m[copyArr[i]] = rank
			rank++
		}
	}
	fmt.Println(m)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], m[arr[i]])
		copyArr[i] = m[arr[i]]
	}
	return copyArr
}
