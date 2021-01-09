package main

import "fmt"

func main() {

	fmt.Println(climbStairs(5, []int{1, 2}))

}

func climbStairs(n int, steps []int) int {
	if n <= 1 {
		return 1
	}

	res := make([]int, n+1)
	res[0] = 1
	res[1] = 1

	total := 0

	for i := 2; i <= n; i++ {
		total = 0
		for _, j := range steps {
			if i-j >= 0 {
				total += res[i-j]
			}
		}
		res[i] = total
	}
	return res[n]
}
