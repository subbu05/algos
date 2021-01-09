package main

import "fmt"

func main() {
	fmt.Println(checkPatternBruteForce("pstring","qst"))
}


// TC O(n*m) SC O(1)
func checkPatternBruteForce(T, P string) int {
	n:=len(T)
	m:=len(P)
	j:=0
	for i:=0;i<=n-m;i++{
		j=0
		for j < m && P[j] == T[i+j] {
			j++
		}
		if j == m {
			return i
		}
	}
	return -1
}

// Rabin Karp Method
func checkPatternRK() int {

	return -1
}


// FiniteAutomta
func checkPatternFA() {

}

// KMP Knuth, Morris, Pratt
// TC O(n)
func checkKMP() {

}
