package main

import "fmt"

func main() {
	x := 98
	y := 100

	fmt.Println(x % y)
	fmt.Println(x & (y - 1))

}
