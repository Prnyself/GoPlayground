package main

import (
	"fmt"
	"sync"
	"runtime"
	//"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	_ = sync.WaitGroup{}
	wg := sync.WaitGroup{}
	wg.Add(1000)
	fmt.Println("start:")
	//ch := make(chan int, 1)
	//ch <- 1
	for i:=0; i<1000; i++ {
		//time.Sleep(time.Second)
		//<- ch
		go func(tmp int) {
			defer wg.Done()
			fmt.Println(tmp)
			//ch <- 1
			//wg.Done()
		}(i)
	}
	//<- ch
	wg.Wait()
	//time.Sleep(time.Second)
	fmt.Println("stop!")
}
