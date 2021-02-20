package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/../"))
}
func simplifyPath(path string) string {
	s := []string{}
	strs := strings.Split(path, "/")

	for _, str := range strs {
		fmt.Println(str, len(str))
		if str == "." || str == "" {
			//fmt.Print("f")
			continue
		} else if str == ".." && len(s) > 0 {
			s = s[0 : len(s)-1]
		} else {
			s = append(s, str)
		}
	}
	res := ""
	for _, v := range s {
		res += "/" + v
	}
	if len(res) == 0 {
		return "/"
	}
	return res
}
