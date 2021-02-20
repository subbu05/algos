package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(coinChange([]int{1}, 5))
}
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Ints(coins)
	c := 0
	for i := len(coins) - 1; i >= 0; i-- {
		fmt.Println(amount / coins[i])
		c += amount / coins[i]
		amount = amount % coins[i]
	}
	return c
}
