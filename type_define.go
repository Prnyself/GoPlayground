package main

import (
	"fmt"
	"reflect"
)

type myint int

type myinteq = int

func main() {
	var i myint = 1
	var ie myinteq = 1
	ShowInt(ie)
	i.ShowInt()
	fmt.Printf("i kind: %v, ie kind: %v\n", reflect.TypeOf(i).Kind(), reflect.TypeOf(ie).Kind())
}

func ShowInt(i int) {
	fmt.Println(i)
}

func (mi myint) ShowInt() {
	fmt.Println(mi)
}
