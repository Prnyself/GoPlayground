package main

import (
	"fmt"
	"reflect"
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
	// panic("aaaa")
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
	if err := Func3(); err != nil {
		fmt.Printf("error get in func 2: %v", err)
	}
}

func Func3() (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
			fmt.Println("err in func3:", err)
			fmt.Println("recover in func3")
		}
	}()
	fmt.Println("in Func3")
	i := 1
	reflect.ValueOf(i).Elem()
	return nil
}
