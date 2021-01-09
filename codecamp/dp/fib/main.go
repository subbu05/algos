package main

import "fmt"

func main() {
	fmt.Println(fib1(10))
	fmt.Println(fib2(10))
	fmt.Println(fib3(10))
}

func fib1(n int) int {
	if n <= 2 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

func fib2(n int) int {
	res := make([]int, n+1)
	res[1] = 1
	res[2] = 1
	for i := 3; i <= n; i++ {
		res[i] = res[i-2] + res[i-1]
	}
	return res[n]
}

func fib3(n int) int {
	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
