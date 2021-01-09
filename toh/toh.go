package main

import "fmt"

func main() {
	toh("Blr", "SVL", "DBI", 5)
}

func toh(from, to, aux string, n int) {
	if n == 1 {
		fmt.Printf("Move gem from %s to %s\n", from, to)
		return
	}
	toh(from, aux, to, n-1)
	fmt.Printf("Move gem from %s to %s\n", aux, to)
	toh(aux, to, from, n-1)
}
