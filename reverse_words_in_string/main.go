package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("      hello! string "))
}

func reverseWords(s string) string {
	i := 0
	j := len(s) - 1
	for {
		if s[i] == 32 {
			i++
		}
		if s[j] == 32 {
			j--
		}
		if s[i] != 32 && s[j] != 32 {
			break
		}
	}
	var sb strings.Builder
	sb.WriteString("")
	strs := strings.Split(s[i:j+1], " ")
	k := len(strs) - 1
	for ; k > 0; k-- {
		sb.WriteString(strs[k])
		sb.WriteRune(32)
	}
	sb.WriteString(strs[k])
	return sb.String()
}
