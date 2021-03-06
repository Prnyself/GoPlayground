package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()

	for i := 0; i < 10; i++ {
		go func() {
			c <- i
			time.Sleep(time.Second * time.Duration(i))
		}()
	}

	<-o
}
