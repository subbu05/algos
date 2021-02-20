package main

import "fmt"

func main() {
	fmt.Print(letterCombinations("23"))
}
func letterCombinations(digits string) []string {
	res := []string{}

	m := make(map[string][]string, 0)

	//m["2"]=  []string{"a","b","c"}
	j := 2
	for i := 97; i <= 119; i += 3 {
		m[string(rune(j))] = []string{string(rune(i)), string(rune(i + 1)), string(rune(i + 2))}
		j++
	}
	fmt.Printf("%#v", m)

	return res
}
