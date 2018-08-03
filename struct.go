package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHi1() {
	p.Name = "leon1"

}

func (p Person) SayHi2() {
	p.Name = "leon2"

}

func main() {
	p1 := &Person{Name: "test", Age: 10}
	fmt.Println("name1 : " + p1.Name)
	p1.SayHi1()
	fmt.Println("name2 : " + p1.Name)

	p2 := Person{Name: "test1", Age: 11}
	fmt.Println("name3 : " + p2.Name)
	p2.SayHi2()
	fmt.Println("name4 : " + p2.Name)

	fmt.Println("p1 : ", p1)
	fmt.Println("p2 : ", p2)

}

type A struct{
	A1 int
	B1 string
}
