package main

import "fmt"

func main() {
	fmt.Println(firstNonRepeatingChar("aabddcbe"))
}

func firstNonRepeatingChar(s string) byte {
	repeated := 0
	i := 0
	for i = 0; i < len(s); i++ {
		repeated = 0
		for j := 0; j < len(s); j++ {
			if i != j && s[i] == s[j] {
				repeated = 1
				break
			}
		}

		if repeated == 0 {
			return s[i]
		}
	}
	return 0
}
