package main

import "fmt"

func main1() {
	fmt.Println(lengthOfLongestSubstring("abcbbabcbb"))
}
func lengthOfLongestSubstring(s string) int {
	res := 0
	index := make([]int, 26)
	i := 0
	for j := 0; j < len(s); j++ {
		fmt.Println(index)
		i = max(i, index[int(s[j])-'a'])
		fmt.Println(i)
		res = max(res, j-i+1)
		index[s[j]-'a'] = j + 1
		fmt.Println(index)
		fmt.Println()

	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
