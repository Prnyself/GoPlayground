package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})
	for i := 0; i < 5; i++ {
		go do2(c)
	}
	time.Sleep(2 * time.Second)
	//close(c)
	//time.Sleep(2 * time.Second)
}

func do2(c <-chan struct{}) {
	// 会阻塞直到收到 close
	<-c
	fmt.Println("hello")
}
