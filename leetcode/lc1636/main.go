package main

import (
	"fmt"
	"sort"
)

func main() {
	frequencySort([]int{3, 3, 3, 2, 2, 1, 0})
}

func frequencySort(nums []int) []int {
	tmp := make([]int, 201)
	for _, v := range nums {
		tmp[v+100]++
	}
	sort.Slice(nums, func(a, b int) bool {
		fmt.Println(nums)
		fmt.Println(tmp)
		if tmp[nums[a]+100] == tmp[nums[b]+100] {
			return nums[a] > nums[b]
		}
		return tmp[nums[a]+100] < tmp[nums[b]+100]
	})
	return nums
}
