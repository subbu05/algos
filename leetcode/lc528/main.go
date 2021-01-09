package main

import (
	"fmt"
)

type Solution struct {
	sum int
	w []int
}


func Constructor(w []int) Solution {
	s := Solution{
		w : make([]int,len(w)),
	}
	copy(s.w,w)
	for i:=0;i<len(s.w);i++{
		s.sum+=w[i]
	}
	return s
}


func (this *Solution) PickIndex() int {
	fmt.Print(len(this.w))
	return 1
	//return this.w[rand.Intn(len(this.w))]/this.sum
}

func main() {
	obj:=Constructor([]int{1,3})
	fmt.Print(obj.PickIndex())
}
