package main

import "fmt"

func main() {
	fmt.Println(halvesAreAlike("bookes"))
}

func halvesAreAlike(s string) bool {
	c:=0

	mid:=len(s)/2
	m := make(map[byte]int)

	m['a']=97
	m['e']=101
	m['i']=105
	m['o']=111
	m['u']=117

	fmt.Println(m)
	for i:=0;i<mid;i++{
		if s[i]>=65&&s[i]<=90 {
			//fmt.Println(s[i]+32)
			if _, ok := m[s[i]+32];ok {

				c++
			}
		} else if s[i]>=97 && s[i]<=122 {
			if _, ok := m[s[i]];ok {
				c++
			}
		}
	}

	for i:=mid;i<len(s);i++{
		if s[i]>=65&&s[i]<=90 {
			if _, ok := m[s[i]+32];ok {
				c--
			}
		} else if s[i]>=97 && s[i]<=122 {
			if _, ok := m[s[i]];ok {
				//fmt.Println(m[s[i]+32])
				c--
			}
		}
	}

	return c == 0
}