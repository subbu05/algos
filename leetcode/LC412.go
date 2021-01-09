package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(1))
}

func fizzBuzz(n int) []string {
	res := []string{}

	for i := 1; i <= n; i++ {
		fmt.Println(i)
		if i%3 == 0 && i%5 != 0 {
			fmt.Println(1)
			res = append(res, "Fizz")
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Println(2)
			res = append(res, "Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println(3)
			res = append(res, "FizzBuzz")
		} else {
			fmt.Println(4)
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}
