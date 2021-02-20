package main

import "fmt"

func main() {
	fmt.Print(getSubString([]byte{'a'}, "a"))
}

func getSubString(a []byte, s string) string {
	start := 0
	end := 0
	if len(a) == 0 {
		return ""
	}
	m := make(map[byte]int)
	for i := 0; i < len(a); i++ {
		m[byte(i)]++
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		if _, ok := m[s[i]]; ok {
			if len(m) == (end - start) {
				return string(s[start : end+1])
			}
			end++
		} else {
			start++
		}
	}
	if len(m) == (end - start) {
		fmt.Print("Success")
		return string(s[start:end])
	}

	return ""
}
