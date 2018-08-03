package main

import (
	"fmt"
	"strconv"
)

func main()  {
	stu := Student{"Tom", 24}
	fmt.Println(stu)
}

type Student struct {
	name string
	age int
}

func (stu Student) String() string {
	return fmt.Sprintf("Hi, my name is: %s, and i'm %d years old", stu.name, stu.age)
}

// 如果实现了Error() string方法，则String()方法会被覆盖
func (stu Student) Error() string {
	strconv.Itoa()
	return fmt.Sprintf("error from %s", stu.name)
}