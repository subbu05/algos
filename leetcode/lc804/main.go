package main

import (
	"fmt"
)

func main() {

	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}

func uniqueMorseRepresentations(words []string) int {
	codes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]int, 0)

	for i := 0; i < len(words); i++ {
		var s string = ""
		for _, v := range words[i] {
			//fmt.Println(v-97)
			//os.Exit(0)
			s = s + (codes[v-97])
		}
		m[s] += 1
	}

	fmt.Println(m)
	return len(m)
}
