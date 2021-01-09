package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4, 5}))
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= r
		r *= nums[i]
	}
	return res
}
