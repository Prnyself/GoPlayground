package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	Timeout(timeoutTest, 5)

	fmt.Println()
}

func timeoutTest(t int) (string, int) {
	time.Sleep(time.Duration(t) * time.Second)
	return "aaa", 1
}

func Timeout(f func(a ...interface{}) interface{}, t int) error {
	timeoutChan := make(chan int, 1)
	resChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Duration(t) * time.Second)
		timeoutChan <- 1
	}()
	go func() {
		f()
		resChan <- 1
	}()
	select {
	case <-resChan:
		return nil
	case <-timeoutChan:
		return errors.New(fmt.Sprintf("timeout for %d second", t))
	}
}
