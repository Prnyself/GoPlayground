package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err in main:", err)
			fmt.Println("recover in main")
		}
	}()
	Func1()
	//panic("aaaa")
}

func Func1() {
	fmt.Println("in Func1")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err in func1:", err)
			fmt.Println("recover in func1")
		}
	}()
	Func2()
}

func Func2() {
	fmt.Println("in Func2")
	Func3()
}

func Func3() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err in func3:", err)
			fmt.Println("recover in func3")
		}
	}()
	fmt.Println("in Func3")
	panic("func3 error")
}
