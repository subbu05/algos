package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = nums[i]
		if len(m) > k {
			delete(m, nums[i-k])
		}
	}
	return false
}

/*
Complexity Analysis
Time complexity :
O(n)
n operations of search, delete and insert, each with constant time complexity.

Space complexity :
O(min(n,k)). The extra space required depends on the number of items stored in the hash table, which is the size of the sliding window,
*/
