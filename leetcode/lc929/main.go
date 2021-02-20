package main

import "fmt"
import "strings"

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"}))
}
func numUniqueEmails(emails []string) int {
	m := make(map[string]int, 0)

	for _, e := range emails {
		i := strings.Index(e, "@")

		part1 := e[0:i]
		part2 := e[i+1:]
		if strings.Contains(part1, "+") {
			part1 = part1[0:strings.Index(e, "+")]
		}
		part1 = strings.ReplaceAll(part1, ".", "")
		fmt.Println(part1, part2)
		m[part1+part2]++
	}
	fmt.Println("%#v", m)
	return len(m)
}

func numUniqueEmails1(emails []string) int {
	m := make(map[string]int, 0)
	visited := false
	for _, e := range emails {
		b := []byte{}
		for i := 0; i < len(e); i++ {
			if e[i] == '@' {
				visited = true
			}
			if e[i] != '.' && e[i] != '+' {
				b = append(b, e[i])
			} else if e[i] == '.' && !visited {
				i++
				visited = true
			} else if e[i] == '+' {
				for ; i < len(e) && e[i] != '@'; i++ {

				}
				b = append(b, []byte(e[i:])...)
			}
		}
		fmt.Println(string(b))
		m[string(b)]++
	}

	return len(m)
}
