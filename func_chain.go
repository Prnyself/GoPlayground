package main

import "fmt"

type Fn func(x, y int) int

func (fn Fn) Chain(f Fn) Fn {
	return func(x, y int) int {
		fmt.Println(fn(x, y))
		return f(x, y)
	}
}
func add(x, y int) int {
	fmt.Printf("%d + %d = ", x, y)
	return x + y
}
func minus(x, y int) int {
	fmt.Printf("%d - %d = ", x, y)
	return x - y
}
func mul(x, y int) int {
	fmt.Printf("%d * %d = ", x, y)
	return x * y
}
func divide(x, y int) int {
	fmt.Printf("%d / %d = ", x, y)
	return x / y
}

// implements from the link below:
// https://colobu.com/2019/08/21/decorator-pattern-pipeline-pattern-and-go-web-middlewares/
func main() {
	var result = Fn(add).Chain(Fn(minus)).Chain(Fn(mul)).Chain(Fn(divide))(10, 5)
	fmt.Println(result)
}
