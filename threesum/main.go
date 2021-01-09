package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{4, 3, 0, 1, -1, -2, 2}))
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i-1 >= 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			switch {
			case sum < 0:
				j++
			case sum > 0:
				k--
			case sum == 0:
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j = j + 1; j < k && nums[j] == nums[j-1]; {
					j++
				}
			}
		}
	}

	return res
}
