package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRL"))
}

func balancedStringSplit(s string) int {
	l := 0
	r := 0
	c := 0
	for _, v := range s {
		fmt.Println(v)
		if v == 76 {
			fmt.Println(v)
			l++
		}
		if v == 82 {
			fmt.Println(v)
			r++
		}
		if l == r {
			c++
			l = 0
			r = 0
		}
	}
	return c
}
