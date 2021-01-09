package main

import "fmt"

func main() {
	insertionsort2([]int{5, 4, 1, 3, 2})
	is2()
}

func insertionsort2(a []int) {
	j := 0
	temp := 0
	for i := 1; i < len(a); i++ {
		temp = a[i]
		for j = i - 1; j >= 0 && a[j] > temp; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = temp
	}
	fmt.Println(a)
}

/*func insertionsort(a []int) {
	//fmt.Println(a)
	j:=0
	for i:=1;i<len(a);i++{
		temp:=a[i]
		for j=i-1;j>=0 && a[j] > temp;j--{
			//fmt.Println(a[j])
			a[j+1]=a[j]
		}
		a[j+1]=temp
	}
	fmt.Println(a)
}*/

func InsertionSort2() {
	a := []int{5, 1, 3, 2, 0}
	j := 0
	temp := 0
	for i := 1; i < len(a); i++ {
		temp = a[i]
		j = i
		for ; j > 0 && a[j-1] > temp; j-- {
			a[j] = a[j-1]
		}
		a[j] = temp
	}

	fmt.Println(a)
}
