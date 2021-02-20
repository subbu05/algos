package main

import "fmt"

type MovingAverage struct {
	elm []int
	max int
}

/** Initialize your data structure here. */
func constructor(size int) MovingAverage {
	return MovingAverage{
		elm: []int{},
		max: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	//fmt.Println((this.elm))

	if len(this.elm) == 0 {
		this.elm = append(this.elm, val)
		return float64(val)
	}

	sum := 0
	if len(this.elm) < this.max {
		this.elm = append(this.elm, val)
		//fmt.Println(this.elm)
		for i := 0; i < len(this.elm); i++ {
			sum += this.elm[i]
		}

	} else {
		this.elm = this.elm[1:]
		this.elm = append(this.elm, val)
		for i := 0; i < len(this.elm); i++ {
			sum += this.elm[i]
		}
	}
	return float64(sum) / float64(len(this.elm))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 */
func main() {
	obj := constructor(3)
	fmt.Printf("Result %f\n", obj.Next(1))
	fmt.Printf("Result %f\n", obj.Next(10))
	fmt.Printf("result %f\n", obj.Next(3))
	fmt.Printf("Result %f\n", obj.Next(5))
}
