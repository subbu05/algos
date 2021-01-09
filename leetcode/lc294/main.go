package main

import (
	"fmt"
)

func main() {
	fmt.Println(FlipGame("++++"))
}

func FlipGame(s string) bool {
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i-1] == '+' && s[i] == '+' {
			if len(s[:i-1]) == 1 || len(s[i+1:]) == 1 {
				return true
			}
		}
	}
	return false
}
