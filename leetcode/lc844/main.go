package main

import "fmt"

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
}

func backspaceCompare(S string, T string) bool {
	s_pointer := len(S) - 1
	t_pointer := len(T) - 1
	s_skips := 0
	t_skips := 0

	for s_pointer >= 0 || t_pointer >= 0 {
		for s_pointer >= 0 {
			if S[s_pointer] == '#' {
				s_skips++
				s_pointer--
			} else if s_skips > 0 {
				s_skips--
				s_pointer--
			} else {
				break
			}
		}
		for t_pointer >= 0 {
			if T[t_pointer] == '#' {
				t_skips++
				t_pointer--
			} else if t_skips > 0 {
				t_skips--
				t_pointer--
			} else {
				break
			}

		}

		if s_pointer >= 0 && t_pointer >= 0 && S[s_pointer] != T[t_pointer] {
			return false
		}

		if (s_pointer >= 0) != (t_pointer >= 0) {
			return false
		}
		s_pointer--
		t_pointer--
	}

	return true
}
