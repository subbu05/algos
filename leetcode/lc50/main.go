package main

import "fmt"

func main() {
	fmt.Print(myPowRecursive(5, 3))
	fmt.Println(myPowIterative(-10, 3))
}

//TC O(lg(n)) SC O(1)
func myPowIterative(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	if n < 0 {
		n *= -1
		x = 1 / x
	}

	cp := x
	ans := 1.0
	for i := n; i > 0; i /= 2 {
		if i%2 == 1 {
			ans = ans * cp
		}
		cp *= cp
	}
	return ans
}

// TC,SC O(lg(n))
func myPowRecursive(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	if n < 0 {
		x = 1 / x
		n *= -1
	}

	temp := myPowRecursive(x, n/2)

	if n%2 == 0 {
		return temp * temp
	}

	return temp * temp * x
}
