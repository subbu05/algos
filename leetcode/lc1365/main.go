package main

import "fmt"

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	counts := make([]int, 101)

	for i := 0; i < len(nums); i++ {
		counts[nums[i]] += 1
	}

	fmt.Println(counts)

	for i := 1; i < 101; i++ {
		counts[i] += counts[i-1]
	}

	fmt.Println(counts)

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			res[i] = 0
		} else {
			res[i] = counts[nums[i]-1]
		}
	}

	return res
}
