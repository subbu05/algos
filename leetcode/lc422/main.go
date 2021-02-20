package main

import "fmt"

func main() {
	fmt.Println(validWordSquare([]string{"abc", "bbc"}))
}

func validWordSquare(words []string) bool {
	if len(words) == 0 {
		return true
	}

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			// check if words is more than row length, word small than i, word char not equal
			if j >= len(words) || len(words[j]) <= i || words[j][i] != words[i][j] {
				return false
			}

		}
	}

	return true
}
