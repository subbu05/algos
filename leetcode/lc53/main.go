//Leetcode 53 Maximum Subarray
package main

import "fmt"

// TC O(n) SC O(1)
func greedyAlgo(nums []int) int {
	sum := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		sum = getMax(nums[i], sum+nums[i])
		max = getMax(sum, max)
	}
	return max
}

// TC O(n) SC O(1)
func kadaneAlgo(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		max = getMax(max, nums[i])
	}
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(greedyAlgo([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(kadaneAlgo([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
