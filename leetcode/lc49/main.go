package main

import "fmt"

func main() {
	fmt.Println(groupAnagram([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagram(strs []string) [][]string {
	res := [][]string{}
	m := make(map[string]string)

	for _, v := range strs {
		a := make(map[rune]int, 26)
		for _, s := range v {
			a[s] += 1
		}
		st := ""
		for i := 'a'; i <= 'z'; i++ {
			st += string(a[i]) + string(i)
		}
		fmt.Println(a)
		if _, ok := m[st]; ok {
			m[st] += "," + v
		} else {
			m[st] = v
		}
	}
	for _, v := range m {
		res = append(res, []string{v})
	}
	return res
}
