package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}
	visited := make(map[int]bool, len(nums))

	dfs(&res, nums, []int{}, visited)
	return res

}

func dfs(res *[][]int, nums []int, cur []int, visited map[int]bool) {
	if len(nums) == len(cur) {
		*res = append(*res, cur)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] == true {
			continue
		}
		visited[i] = true
		dfs(res, nums, append(cur, nums[i]), visited)
		visited[i] = false
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
