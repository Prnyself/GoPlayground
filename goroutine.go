package main

import "fmt"

func main() {
	c := make(chan int, 1) //修改2为1就报错，修改2为3可以正常运行
	for {
		go func() { c <- 1 }()
		go func() { c <- 3 }()
		go func() { c <- 2 }()
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c)
	}
}
