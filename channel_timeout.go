package main

import (
	"fmt"
	"time"
)

func main() {
	done := do()
	select {
	case <-done:
		fmt.Println("done")
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}
}

func do() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 4)
		done <- struct{}{}
	}()
	fmt.Println("done in do")
	return done
}
