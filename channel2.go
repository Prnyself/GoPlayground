package main

import (
	"fmt"
	//"time"
)

func main() {
	sourceNum := 4
	allJobNum := 20
	chan1 := make(chan int, sourceNum)
	chan2 := make(chan int, allJobNum)
	//var test int
	for i := 0; i < sourceNum; i++ {
		//go func() {
		chan1 <- 1
		//}()
	}

	for i := 0; i < allJobNum; i++ {
		go get(i, chan1, chan2)
	}

	for i := 0; i < allJobNum; i++ {
		<-chan2
	}

	fmt.Println("all work done!")
}

func get(i int, pschan chan int, asynchan chan int) {
	<-pschan
	fmt.Println(i)
	//time.Sleep(time.Second)
	asynchan <- 1
	pschan <- 1
}
