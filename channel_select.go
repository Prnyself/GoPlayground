package main

import "fmt"

func initFunc() (chan int, chan int) {
	chan1 := make(chan int, 10)
	chan2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		chan1 <- 1
		chan2 <- 2
	}
	return chan1, chan2
}

func main() {
	chan1, chan2 := initFunc()
	select {
	case i := <-chan1:
		fmt.Println("chan", i, "selected")
	case i := <-chan2:
		fmt.Println("chan", i, "selected")
	default:
		fmt.Println("default")
	}
}
