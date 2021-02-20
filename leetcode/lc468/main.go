package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(validIPAddress("0:0db8:85a3:0000:0000:8a2e:0370:7334"))
}

func validIPAddress(IP string) string {
	parts := strings.Split(IP, ":")
	if len(parts) == 8 {
		return v6(parts)
	}
	parts = strings.Split(IP, ".")
	if len(parts) == 4 {
		return v4(parts)
	}
	return "Neither"
}

func v6(parts []string) string {
	for _, v := range parts {
		if validHexa(v) == false {
			return "Neither"
		}
		/*i, er := strconv.Atoi(v)
		if er != nil || i > 0xffff || i < 0 {
			return "Neither"
		}*/
	}
	return "IPv6"
}

func v4(parts []string) string {
	// TBD to check "023"
	for _, v := range parts {
		for i := 0; i < len(v)-1; i++ {
			if v[i] == '0' {
				return "Neither"
			} else {
				break
			}
		}
		i, er := strconv.Atoi(v)
		if er != nil || i > 255 || i < 0 {
			return "Neither"
		}
	}
	return "IPv4"
}

func validHexa(s string) bool {
	if len(s) == 0 || len(s) > 4 {
		return false
	}
	for i := 0; i < len(s); i++ {

		if (s[i] >= 48 && s[i] <= 57) || (s[i] >= 65 && s[i] <= 70) || (s[i] >= 97 && s[i] <= 102) {
		} else {
			fmt.Println(s[i])
			return false
		}
	}
	return true
}
