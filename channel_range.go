package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	input := make(chan int)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(input chan<- int, i int) {
			input <- i
			wg.Done()
		}(input, i)
	}
	// if we do not close input, it will panic with "all goroutines are asleep - deadlock!"
	go func() {
		wg.Wait()
		// fmt.Println("should close input here")
		close(input)
	}()

	for v := range input {
		fmt.Println("v:", v)
	}
	fmt.Println("active goroutine:", runtime.NumGoroutine())
}
