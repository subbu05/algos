package main

import "fmt"

func main() {
	fmt.Println(halvesAreAlike("Uo"))
}
func halvesAreAlike(s string) bool {
	r := []rune(s)
	return check(r[0:len(r)/2]) == check(r[len(r)/2:])
}

func check(r []rune) int {
	c := 0
	for i := 0; i < len(r); i++ {
		if r[i] >= 65 && r[i] <= 90 {
			r[i] += 32
		}
		if r[i] == 'a' || r[i] == 'e' || r[i] == 'o' || r[i] == 'u' || r[i] == 'i' {
			c++
		}
	}
	fmt.Println(c)

	return c
}
