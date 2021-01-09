package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 5
	}()

	time.Sleep(100)
	fmt.Println(<-ch)

}
