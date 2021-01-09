package main

import "fmt"

func main() {
	fmt.Print(minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}}))
}

func minTimeToVisitAllPoints(points [][]int) int {
	sum3 := 0
	sum1 := 0
	sum2 := 0
	for i := 0; i < len(points)-1; i++ {
		sum1 = abs(points[i][0] - points[i+1][0])
		sum2 = abs(points[i][1] - points[i+1][1])
		if sum1 > sum2 {
			sum3 += sum1
		} else {
			sum3 += sum2
		}
	}
	return sum3
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
