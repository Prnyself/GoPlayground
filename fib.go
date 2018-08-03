package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		index, res := fib(i)
		fmt.Println("index:", index, "res:", res)
	}
	recur(10)

	fv := func() {
		fmt.Println("hello, world")
	}
	fv()
	fmt.Printf("type: %T\n", fv)
	fmt.Println(f())

	newFib := fibNew()
	for i := 0; i <= 10; i++ {
		fmt.Println("fib index is", i, "value is", newFib(i))
	}

	//f := func(x, y int) int {
	//	return x + y
	//}
	//fmt.Println(f(4,6))
}

func fib(i int) (index, res int) {
	if i <= 1 {
		index = i
		res = 1
	} else {
		index = i
		_, res1 := fib(i - 1)
		_, res2 := fib(i - 2)
		res = res1 + res2
	}
	return
}

func recur(i int) {
	fmt.Println(i)
	if i <= 1 {
		return
	}
	recur(i - 1)
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func fibNew() func(int) int {
	var index1, index2 int
	return func(i int) int {
		if i <= 1 {
			index1 = i
			index2 = 1
		} else {
			a := index1
			index1 = index2
			index2 = a + index2
		}
		fmt.Println("index1:", index1, "index2:", index2)
		return index2
	}
}
