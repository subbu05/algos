package main

import (
	"fmt"
)

func main() {
	fmt.Println(nthUglyNumber(25))
}

func nthUglyNumber(n int) int {
	res := make([]int, n) // to store the ith ugly number
	res[0] = 1            // first ugly number 1
	a, b, c := 0, 0, 0
	f2, f3, f5 := 0, 0, 0
	for i := 1; i < len(res); i++ {
		f2 = res[a] * 2
		f3 = res[b] * 3
		f5 = res[c] * 5
		res[i] = min(f2, f3, f5)
		if res[i] == f2 {
			a++
		}
		if res[i] == f3 {
			b++
		}
		if res[i] == f5 {
			c++
		}
	}
	return res[n-1]
}

func min(args ...int) int {
	m := 1<<31 - 1
	for _, v := range args {
		if v < m {
			m = v
		}
	}
	return m
}

func isUgly(num int) bool {
	for num <= 0 {
		return false
	}

	for num%30 == 0 {
		num /= 30
	}

	for num%2 == 0 {
		num /= 2
	}

	for num%3 == 0 {
		num /= 3
	}

	for num%5 == 0 {
		num /= 5
	}

	return num == 1
}
