package main

import "fmt"

var a string

type myInt int

func main() {
	a = "G"
	print(a) // G
	f1()
	//s1 := new(MyImpStruct)
	s1 := &MyImpStruct{}
	s1.Name = "test_Name"
	s1.start()
	fmt.Println(s1.Name)

	foo := myInt(100)
	foo.Increase()
	fmt.Println(foo)
}

func f1() {
	a := "O"
	print(a) // 0
	f2()
}

func f2() {
	print(a) // G
}

type MyStruct struct {
	Name string
}

type MyImpStruct struct {
	MyStruct
}

func (s1 *MyImpStruct)start()  {
	t := &MyStruct{Name:s1.Name}
	t2 := &MyImpStruct{MyStruct: *t}
	str := s1.Name
	fmt.Println(str)
	fmt.Println(t2.MyStruct.Name)
	s1.Name = "modifiedName"
}

func (foo *myInt)Increase(){
	*foo = *foo + 100
}