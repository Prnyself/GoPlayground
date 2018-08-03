package main

import "fmt"

func main() {
	mapTest := make(map[int]string, 10)
	mapTest[1] = "alpha"
	mapTest[3] = "alpha3"

	for i := 1; i <= 3; i++ {
		value, ok := mapTest[i]
		if !ok {
			fmt.Println("not found at", i)
		} else {
			fmt.Println("value at", i, "is", value)
		}
	}

	for k, v := range mapTest {
		fmt.Println(k, "===", v)
	}

	//fmt.Printf()
	//fmt.Sprintf()
	//fmt.Fprintf()
}
