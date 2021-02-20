package main

import "fmt"

func main() {
	fmt.Println(canPermutePalindrome("carerac"))
}
func canPermutePalindrome(s string) bool {
	m := make(map[rune]int)
	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v] -= 1
		} else {
			m[v] += 1
		}
	}
	odds := 0
	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v)
		if v%2 != 0 {
			odds++
		}

		fmt.Println(odds)
	}

	if len(m) <= 1 {
		return true
	}
	return (odds >= 2) == false
}
