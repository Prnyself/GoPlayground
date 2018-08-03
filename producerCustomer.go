package main

import (
	"time"
	"fmt"
)

func main() {
	channel := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go producer(channel, i)
	}
	for i := 0; i < 12; i++ {
		go consumer(channel, i)
	}
	for {
		time.Sleep(time.Second)
	}

}

//生产者 只可以进行发送操作
func producer(channel chan<- int, i int) {
	//for i := 0; i < 101; i++ {
		channel <- i
		i %= 100
		time.Sleep(time.Millisecond * 100)
	//}
}

//生产者 只可以进行接收操作
func consumer(channel <-chan int, id int) {
	for {
		v := <-channel
		fmt.Println(id, v)
		time.Sleep(time.Millisecond * 200)
	}
}
