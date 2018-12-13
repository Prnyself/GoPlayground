package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1) //First
	exit := make(chan int)
	go func() {
		fmt.Println("goroutine")
		close(exit)
		for {
			if true {
				//fmt.Println("Looping!")  //Second
			}
		}
	}()
	fmt.Println(<-exit)
	fmt.Println("Am I printed?")
}
