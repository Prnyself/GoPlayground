package main

import(
	"fmt"
	"reflect"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	d := make(chan bool)
	fmt.Println(d)
	fmt.Println("hello this is playground")
	var p *int
	a := 100
	p = &a
	c := &a
	fmt.Println(p)
	fmt.Println(c)
	//fmt.Println(make(chan bool))

	modify(&a)
	fmt.Println(a)

	a = modify2(a)
	fmt.Println(a)

	v := Vertex{1, 2}
	newPointer := &v
	(*newPointer).Y = 4
	newPointer.X = 1e9 //其实是应该写成(*p).X的，但是go支持隐式间接引用，直接写p.X就可以
	fmt.Println(v)
}

func modify(a *int) {
	*a = *a + 1
}

func modify2(a int) int{
	a = a + 1
	return a
}
