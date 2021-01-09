// https://leetcode.com/problems/arranging-coins
package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(ArrangeCoins_N_1(5))

	fmt.Println(ArrangeCoins_LogN_1(5))

	fmt.Println(arrangeCoins_1_1(5))
}

// Bruteforce building ladder from start till end.
func ArrangeCoins_N_1(n int) int {
	c := 0
	i := 1
	for {
		if n >= i {
			n -= i
			i += 1
			c++
		} else {
			break
		}
	}

	return c
}

// using binary search
func ArrangeCoins_LogN_1(n int) int {
	l := 0
	r := n
	mid := 0
	temp := 0
	for l <= r {
		mid = l + (r-l)/2
		temp = mid * (mid + 1) / 2
		if temp == n {
			return mid
		}
		if n < temp {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}

// using formula k(k+1)≤2N
func arrangeCoins_1_1(n int) int {
	return int(math.Sqrt(float64(2*n)+0.25) - 0.5)
}
