package main

import "fmt"

func main() {
	fmt.Print(rotateString("abc", "cab"))
}

func rotateString(A string, B string) bool {
	r := []rune(A)
	for i := 0; i < len(r); i++ {
		fmt.Println(string(r))
		r = append(r[1:], r[0])
		fmt.Println(string(r))
		if string(r) == B {
			return true
		}
	}
	return false
}
