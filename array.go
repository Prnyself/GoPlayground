package main

import "fmt"

func main() {
	arr1 := [5]int{1,2,3,4,5}
	arr2 := arr1
	arr3 := &arr1

	fmt.Println("arr1: ", &arr1[0])
	fmt.Println("arr1(&): ", &arr1)
	fmt.Println("arr2: ", &arr2[0])
	fmt.Println("*arr3: ", &(*arr3)[0])
}
