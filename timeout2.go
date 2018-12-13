package main

import (
	"fmt"
	"time"
)

func main() {
	var e int
	ok := true
	var timer *time.Timer
	ch11 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(i) * time.Second)
			ch11 <- i
		}
	}()
	for {
		select {
		case e = <-ch11:
			fmt.Println("ch11 -> ", e)
		case <-func() <-chan time.Time {
			if timer == nil {
				timer = time.NewTimer(3 * time.Second)
			} else {
				timer.Reset(3 * time.Second)
			}
			return timer.C
		}():
			fmt.Println("Timeout")
			ok = false
			break
		}
		if !ok {
			break
		}
	}

}
