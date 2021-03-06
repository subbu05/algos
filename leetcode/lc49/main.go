package main

import "fmt"

func main() {
	fmt.Println(groupAnagrams2([]string{"aet", "tea", "tan", "ate", "nat", "bat"}))
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
		//fmt.Println(a)
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

func groupAnagrams2(strs []string) [][]string {
	res := [][]string{}
	m := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		c := make([]rune, 26)
		r := []rune(strs[i])
		for j := 0; j < len(r); j++ {
			c[r[j]-'a']++
		}
		//fmt.Println(c)
		str := ""
		for k := 0; k < 26; k++ {
			str = str + "!" + string(c[k])
		}
		fmt.Println(str)
		m[str] = append(m[str], strs[i])
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
