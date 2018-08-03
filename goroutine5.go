package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	for i := 0; i < 100000; i++ {
		go fmt.Println("now i =", i)
		//fmt.Println("now i =", i)
		//time.Sleep(time.Second * 3)
	}
	now_end := time.Now()
	fmt.Println("time use:", now_end.Sub(now))

}
