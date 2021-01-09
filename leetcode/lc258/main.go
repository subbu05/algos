package main

import "fmt"

func main() {
	fmt.Println(addDigits(23))
}

func addDigits(num int) int {
	if num <= 9 {
		return num
	}
	return 1 + (num-1)%9
}
