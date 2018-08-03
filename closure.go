package main

import "fmt"

func main() {

	add := func(x int) (func(int)) {
		i, j := 0, x

		return func(xx int) {
			i++
			j = xx
			fmt.Println("now i =", i)
			fmt.Println("now j =", j)
		}
	}

	ad := add(0)

	add(0)(3)
	add(5)(3)
	add(3)(3)
	ad(2)
	ad(3)
	ad(1)
	AddTest()
	AddTest()
	AddTest()
	AddTest()
}

func AddTest() {
	i := 0
	i++
	fmt.Println("i in func is:", i)
}
