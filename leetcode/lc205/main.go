package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("abb", "egg"))
}

func isIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		s1, ok1 := m1[s[i]]
		t1, ok2 := m2[t[i]]

		if (ok1 == false && ok2 == true) || (ok1 == true && ok2 == false) {
			return false
		}
		if ok1 == false && ok2 == false {
			m1[s[i]] = t1
			m2[t[i]] = s1
		} else if s[i] != t1 && t[i] != s1 {
			return false
		}
	}
	return true
}

func isIsomorphic2(s, t string) bool {
	m := make([]int, 512)
	for i := 0; i < len(s); i++ {
		if m[s[i]] != m[int(t[i])+256] {
			return false
		}
		m[s[i]] = m[int(t[i])+256]
		m[int(t[i])+256] = i + 1
	}
	return true
}
