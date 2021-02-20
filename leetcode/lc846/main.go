package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
}

func isNStraightHand(hand []int, W int) bool {
	m := make(map[int]int)
	for _, v := range hand {
		m[v]++
	}

	keys := make([]int, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Printf("%#v\n", m)

	for len(m) > 0 {
		first := keys[0]
		fmt.Println(keys)
		for i := first; i < first+W; i++ {
			fmt.Println(i)
			v, ok := m[i]
			if ok == false {
				return false
			}
			if v == 1 {
				delete(m, i)
			} else {
				m[i]--
			}
		}
	}
	return true
}
