package main

import "fmt"

func main() {
	fmt.Print(partitionLabels("ababcbacadefegdehijhklij"))
}

func partitionLabels(S string) []int {
	last := make([]int, 26)
	for i := 0; i < len(S); i++ {
		last[S[i]-'a'] = i
	}
	fmt.Println(last)
	j := 0
	anchor := 0
	res := []int{}
	for i := 0; i < len(S); i++ {
		if j > last[S[i]-'a'] {
			j = last[S[i]-'a']
		}
		if i == j {
			res = append(res, i-anchor+1)
			anchor = i + 1
		}
	}
	return res
}
