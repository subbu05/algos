package lc807

import "fmt"

func main() {
	fmt.Println(maxIncreaseKeepingSkyline([][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	r := make([]int, len(grid))
	c := make([]int, len(grid))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			r[i] = max(r[i], grid[i][j])
			c[j] = max(c[j], grid[i][j])
		}
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			res += min(r[i], c[j]) - grid[i][j]
		}
	}
	return res
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
