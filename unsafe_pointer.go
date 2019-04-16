package main

import (
	"fmt"
	"unsafe"
)

type Vstu struct {
	age      int
	PhoneNum string
	name     string
}

func main() {
	v := &Vstu{PhoneNum: "150"}
	va := (*int)(unsafe.Pointer(v))
	*va = 28
	fmt.Println("v after modify age", v)
	vn := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + unsafe.Offsetof(Vstu{}.name)))
	//vn := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + unsafe.Sizeof(Vstu{}.age) + unsafe.Sizeof(Vstu{}.phone)))
	// 上边两行等同，但是offsetof参数只能是struct field，但是可以直接到该字段的开头，更方便
	*vn = "lance"
	fmt.Println("v after modify name", v)
}
