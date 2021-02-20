package main

import "fmt"

func main() {
	fmt.Println(reverseWords("apple car"))
}

func reverseWords(s string) string {
	r := []rune(s)
	reverse(r, 0, len(r)-1)
	fmt.Println(string(r))
	start := 0
	for i := 0; i <= len(r); i++ {
		if i == len(r) || r[i] == ' ' {
			reverse(r, start, i-1)
			fmt.Println(string(r))
			start = i + 1
		}
	}
	return string(r)
}

func reverse(s []rune, st, en int) {
	for ; st < en; st, en = st+1, en-1 {
		s[st], s[en] = s[en], s[st]
	}

}
