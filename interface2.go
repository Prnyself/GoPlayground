package main

import "fmt"

type C struct {
	Name string
	sex  string
}

type E interface {
	click()
}

func (*C) click() {
	fmt.Println("click~!~!~!")
}

func main() {
	//var c = new(C)
	c := new(C)
	var inter E
	inter = c
	Touch(inter)
}

func Touch(e E) {
	e.click()
}
