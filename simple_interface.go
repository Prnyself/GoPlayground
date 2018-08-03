package main

import "fmt"

type ISimple interface {
	Get() int
	Set(value int)
}

type Simple struct {
	value int
}

func (ss *Simple) Get() int {
	return ss.value
}

func (ss *Simple) Set(value int) {
	ss.value = value
}

func main() {
	simple := new(Simple)
	ss := ISimple(simple)
	ss.Set(456)

	meStruct := &Simple{} //返回一个结构体，是类型的
	fmt.Printf("struct type is %T\n", meStruct)
	fmt.Println(meStruct)

	fmt.Println(ss.Get())
	fmt.Println(simple.value)

	var object User
	test:= mycopy(&object)

	fmt.Println(test)
	fmt.Printf("test type is %T\n", test)

	//num := 123
	//var p1 *int
	//var p2 *int
	//p2 = &num
	//p1 = p2
	////*p1 = &num
	//fmt.Println(*p1)

	//ass, ok := ss.(ISimple)
	ass, ok := ss.(ISimple)
	if !ok{
		fmt.Println("type is wrong")
	}
	fmt.Println(ass)
	fmt.Println(ass.Get())
	fmt.Printf("ass type is %T\n", ass)
}

type User struct{
	Name string
}

func mycopy(object interface{}) interface{}{
	var user [2]User
	user[0].Name = "oscar"

	object = &user[0]
	return object

}
