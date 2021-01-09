package main

import "fmt"

func main() {
	a := 1
	b := 1
	for a+b < 4000000 {
		a, b = b, a+b
	}
	fmt.Println(b)
}
