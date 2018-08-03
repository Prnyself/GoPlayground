package main

import (
	"os"
	"fmt"
	"time"
)

func main()  {
	channel1 := make(chan string)
	go writeRoutine(channel1)
	go readRoutine(channel1)

	os.Stdin.Read(make([]byte, 1))
	chanTest := make(chan string)
	chanTest <- "test"
	fmt.Println(<- chanTest)

}

func readRoutine(ch chan string)  {
	for{
		fmt.Printf("message: %s\n", <-ch)
		time.Sleep(time.Duration(5) * time.Second)
	}
}

func writeRoutine(ch chan string)  {
	for{
		ch <- time.Now().String()
		time.Sleep(time.Duration(5) * time.Second)
		//time.Sleep(time.Duration(1000 * 100))
	}
}
