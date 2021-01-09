package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countPrimes2(10))
}

func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}

	notPrimes := make([]bool, n)
	notPrimes[0] = true
	notPrimes[1] = true
	for j := 2; j < int(math.Sqrt(float64(n))); j++ {
		if !notPrimes[j] {
			for k := 2; k*j < n; k++ {
				notPrimes[k*j] = true
			}
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		if notPrimes[i] == true {
			sum++
		}
	}
	return sum
}

func countPrimes2(n int) int {
	if n < 3 {
		return 0
	}
	c := n / 2
	f := make([]bool, n)
	for i := 3; i*i < n; i += 2 {
		if f[i] {
			continue
		}
		for j := i * i; j < n; j += i * 2 {
			if !f[i] {
				c--
				f[i] = true
			}
		}
	}
	return c
}
