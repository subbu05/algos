package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotateArray(numbers, k)
	//fmt.Println(numbers)
	//fmt.Printf("%d",&numbers)
	checkPerfectNumber(28)
}

func rotateArray(numbers []int, k int) {

	a := append(numbers[len(numbers)-k:], numbers[0:len(numbers)-k]...)
	numbers = a
	fmt.Println(numbers)

}

func checkPerfectNumber(num int) bool {
	temp := 1
	for i := 2; i < num; i++ {
		if num%i == 0 {
			temp += num / i
		}
	}
	fmt.Println(temp)
	return temp == num
}
