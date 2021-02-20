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

func canWin(s string) bool {
	if len(s) < 2 {
		return false
	}
	return helper(s, make(map[string]bool))
}

func helper(s string, m map[string]bool) bool {
	if v, ok := m[s]; ok {
		return v
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '+' && s[i+1] == '+' {
			t := s[0:i] + "--" + s[i+2:]
			if !helper(t, m) {
				m[s] = true
				return true
			}
		}
	}

	m[s] = false
	return false
}
