package main

import "fmt"

func main() {
	//a := []int{5,4,3,2,1}
	fmt.Println(aa())
	/*for out:=len(a)-1;out > 0 ;out --{
		for i:=0;i<out;i++{
			if a[i]>a[out]{
				a[i],a[out]=a[out],a[i]
			}
		}
	}*/

	//fmt.Println(a)
}

func aa() []int {
	a := []int{4, 3, 2, 1, 0}
	for out := len(a) - 1; out > 0; out-- {
		for i := 0; i < out; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a
}
