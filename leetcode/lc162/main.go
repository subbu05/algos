package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findPeakElement([]int{1, 2}))
}

func findPeakElement(nums []int) int {
	nums = append([]int{int(math.MinInt64)}, nums...)
	nums = append(nums, -1)
	i := 0
	for i = 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			break
		}
	}
	return i - 1
}
