package main

import (
	"fmt"
)

func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	if n < 1 {
		return 0
	}
	p := make([]bool, n)
	for i := 2; i < n; i++ {
		p[i] = true
	}
	for j := 2; j*j < n; j++ {
		if p[j] == true {
			for k := j * j; k < n; k += j {
				p[k] = false
			}
		}
	}
	c := 0
	for i := 0; i < n; i++ {
		if p[i] == true {
			c++
		}

	}
	return c
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
