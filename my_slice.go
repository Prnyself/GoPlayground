package main

import "fmt"

func printType(t ...interface{}) {
	for _, v := range t {
		switch v.(type) {
		case []int:
			fmt.Println("type is []int, value is", v)
		case int:
			fmt.Println("type is int, value is", v)
		case interface{}:
			fmt.Println("type is interface, value is", v)
		default:
			fmt.Println("type unknown, value is", v)
		}
	}
}

func myPrint(t ...int) {
	for _, v:= range t{
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("start test")
	slice := []int{1, 2, 3, 4, 5}
	myPrint(slice...)
	fmt.Println(slice)

	slicePointer := make([]interface{}, 5)
	for i, v := range slice{
		slicePointer[i] = v
	}
	fmt.Printf("slicePointer type is %T\n", slicePointer[0])

	fmt.Println(slicePointer)
	printType(slice)
	//printType(1,2,3)
	printType(slicePointer...)
}
