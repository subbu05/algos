package main

import "fmt"

func main() {
	fmt.Print(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
}

func createTargetArray(nums []int, index []int) []int {
	res := []int{}

	for i := 0; i < len(nums); i++ {
		ind := index[i]
		res = append(res[0:ind], append([]int{nums[i]}, res[ind:]...)...)
	}
	return res
}
