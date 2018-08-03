package main

import "fmt"

type C struct {
	Name string
	sex  string
}

type D struct {
	Age int
}

//type DD struct {
//	Age int
//}

type E interface {
	click()
}

func (*C) click() {
	fmt.Println("click~!~!~!")
}

func (d *D) click() {
	fmt.Println("knock, ", d.Age)
}

func main() {
	c := new(C)
	var cc E = c
	fmt.Printf("c type is %T\n", c)
	fmt.Printf("cc type is %T\n", cc)
	Touch(c)
	cc.click()
	d := new(D)
	d.Age = 40
	Touch(d)
	math.sq
	//dd := new(DD)
	//var ee E = dd
}

func Touch(e E) {
	e.click()
}
