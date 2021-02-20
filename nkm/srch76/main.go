package main

import "fmt"

func main() {
	fmt.Println(trailingzeros(10))
}

func trailingzeros(n int) int {
	c := 0

	if n < 0 {
		return -1
	}
	temp := 0
	for i := 5; ; i *= 5 {
		temp = n / i
		fmt.Println(temp)
		if temp > 0 {
			c += temp
		} else {
			break
		}
	}
	return c
}
