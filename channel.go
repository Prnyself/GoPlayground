package main

import "fmt"

var i int = 0

func Count(ch chan int) {
	ch <- 1
	i += 1
	fmt.Println("Counting:", "i=", i)
	//在fmt.Println 语句执行的时候 你十条通道的数据都发送并接收完成了，程序也就退出了，
	//多执行几次你会发现你那样写Counting的输出数量不定
}


func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range chs {
		<-ch
	}

	ch1 := make(chan *int, 1)
	tmp := 1
	//fmt.Printf("%p\n", &tmp)
	fmt.Printf("tmp value: %v, addr: %p\n", tmp, &tmp)
	ch1 <- &tmp
	tmp2 := <- ch1
	//fmt.Printf("%p\n", &tmp2)
	tmp = 2
	fmt.Printf("tmp2 value: %v, addr: %p\n", *tmp2, tmp2)
}

