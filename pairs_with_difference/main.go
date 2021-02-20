package main

import "fmt"

func main() {
	fmt.Println(GetResult([]int{0, -1, 1, -2, 2}, 1))
}

func GetResult(nums []int, k int) [][]int {
	m := make(map[int]int, 0)
	res := make([][]int, 0)

	for {
		m[nums[i]-k] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		diff := nums[i] - k
		if ind, ok := m[diff]; ok {
			res = append(res, []int{ind, i})
		} else {
			m[nums[i]] = i
		}
	}
	return res
}
