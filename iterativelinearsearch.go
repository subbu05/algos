package main

import "fmt"

func main() {
	a := []int{5,2,8,2,4,3,9,0,6,1}
	fmt.Println(search(a,5)) // prints true
	fmt.Println(search(a,25)) //prints false
}

// finds the key in the array and returns true if found else false.
func search(a []int, key int) bool {
	for _,v := range a {
		if v == key {
			return true
		}
	}
	return false
}