package main

import "fmt"

func main() {
	i := 0
	t := &i
	defer fmt.Println(*t)  // defer在之后return后执行，但值是在这里确定好的
	defer fmt.Println("second defer") // 栈，先进后出
	fmt.Println("before change: ", *t)
	i = 10
	fmt.Println("after change: ", *t)
	//before change:  0
	//after change:  10
	//0
}

func excep(i int) int{
	a := i
	fmt.Println(a)
	return a
}