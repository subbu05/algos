package main

import "fmt"

func main() {
	var temp = calculateTime("pqrstuvwxyzabcdefghijklmno", "leetcode")
	fmt.Println(temp)
}

func calculateTime(keyboard string, word string) int {
	kb := []byte(keyboard)
	m := make(map[byte]int, 0)
	for _, v := range kb {
		m[v] = int(v - 'a')
	}

	// pc:=0
	sum := m[word[0]]
	//fmt.Println(sum)
	fmt.Println('d')
	for i := 1; i < len(word); i++ {
		sum += absDiff(m[word[i]], m[word[i-1]])
	}
	return sum
}

func absDiff(i, j int) int {
	if i >= j {
		return i - j
	}
	return j - i
}
