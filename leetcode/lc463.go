package main

import "fmt"

func main() {
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 0, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
}

func islandPerimeter(grid [][]int) int {
	ret := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				ret += 4
				if i > 0 && grid[i-1][j] == 1 {
					ret -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					ret -= 2
				}
			}
		}
	}
	return ret
}
