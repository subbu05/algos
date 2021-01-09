package main

import "fmt"

func xorOperation(n int, start int) int {
	res := 0
	i := 0
	for i < n {
		fmt.Println(start + 2*i)
		res ^= (start + 2*i)
		i++
	}
	return res
}

func main() {
	fmt.Println(xorOperation(5, 0))
}
