package main

import (
	"fmt"
)

func main() {
	fmt.Println(selfDividingNumbers(15, 15))
}

func selfDividingNumbers(left int, right int) []int {
	res := []int{}

	for i := left; i <= right; i++ {
		if i <= 9 {
			res = append(res, i)
			continue
		}
		if i%10 == 0 {
			continue
		}
		j := i
		for j > 0 {
			rem := j % 10
			fmt.Printf("%d %d", j, rem)
			if rem == 0 {
				fmt.Println("Continuing")
				continue
			}
			if i%rem == 0 {
				fmt.Println("Line 31")
				j /= 10
			} else {
				fmt.Println("Line 34")
				break
			}

		}
		if j == 0 {
			res = append(res, i)
		}
	}
	return res
}
