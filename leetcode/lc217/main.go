package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 1}))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok == true {
			return true
		} else {
			m[nums[i]] = 1
		}
	}
	return false
}
