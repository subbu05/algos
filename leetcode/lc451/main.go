package main

import (
	"fmt"
	"sort"
	"strings"
)

type rr []rune

func main() {
	fmt.Println(helper("dd"))
}

func helper1(s string) string {
	if len(s) <= 2 {
		return s
	}
	r := []string{}
	t := []string{}

	for i := 0; i < len(s); i++ {
		r = append(r, string(s[i]))
	}
	sort.Strings(r)

	start := 0
	i := 0
	for i = 1; i < len(r); i++ {
		if r[i] != r[i-1] {
			t = append(t, r[start:i]...)
			start = i
		}
	}
	t = append(t, r[start:i]...)
	sort.SliceStable(t, func(i, j int) bool {
		return len(t[i]) < len(t[j])
	})

	res := ""
	for i := 0; i < len(t); i++ {
		res += t[i]
	}
	return res

}

func helper(s string) string {
	if len(s) <= 2 {
		return s
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	keys := []int{}
	for _, k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	res := ""
	for i := len(keys) - 1; i >= 0; i-- {
		for k, v := range m {
			if keys[i] == v {
				res += strings.Repeat(string(k), v)
				delete(m, k)
				break
			}
		}

	}
	return res
}
