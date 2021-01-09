// LC5 Longest Palindromic Substring TC O(n2) SC O(1)
package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("aabbaa"))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	st, end := 0, 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		l1 := helper(s, i, i)
		l2 := helper(s, i, i+1)
		if l1 > l2 {
			maxLen = l1
		} else {
			maxLen = l2
		}
		if maxLen > end-st {
			st = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[st : end+1]
}

func helper(s string, L, R int) int {
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}
