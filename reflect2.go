package main

import (
	"reflect"
	"fmt"
)

func main() {
	var x float64 = 3.4
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind().String() == "float64")
	fmt.Println("value:", v.Float())

	fmt.Println("t type:", t.String())
}
