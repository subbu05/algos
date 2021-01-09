package main

import "strings"

func main() {
	bos(9, 8, 9, "bacbacacb")
}

func bos(n, add, app int, str string) int {
	if n <= 1 || len(str) <= 1 {
		return n
	}
	dstr := ""
	r := []rune(str)
	start := 0
	for i := 0; i < n; i++ {
		temp := strings.Contains(dstr, string(r[start:i]))
		if temp == true {
			continue
		} else {

		}
	}
}
