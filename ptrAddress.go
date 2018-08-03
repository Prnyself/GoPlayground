package main

import (
	"fmt"
)

func main() {
	a := make(map[int]string, 1)
	a[1] = "a"
	fmt.Printf("a: %p\n", a)
	fmt.Printf("&a: %p\n", &a)
	a[3] = "c"
	fmt.Printf("a: %p\n", a)
	fmt.Printf("&a: %p\n", &a) // a的cap为1, 添加a[3]发生扩容, 但是a和&a都没变, 说明扩容不会改变地址

	modMap(a)
	fmt.Println("a value: ", a) // 修改后的a变成{1:a 3:c 5:t}
	fmt.Printf("a after mod: %p\n", a)
	fmt.Printf("&a after mod: %p\n", &a) // a在函数修改之后a和&a也没变

	newA := a
	fmt.Printf("newA: %p\n", newA)  // a赋值给了newA, 地址不变
	fmt.Printf("&newA: %p\n", &newA)  // a和newA两个变量名不同所以 &newA与&a不一样
	newA = map[int]string{11:"aa"}  // newA从新赋值
	fmt.Printf("newA: %p\n", newA)  // newA地址变化
	fmt.Printf("&newA: %p\n", &newA)  // 但&newA不变

	modPtr(&newA)
	fmt.Println("newA after modPtr: ", newA)
	fmt.Printf("newA: %p\n", newA)  // newA地址变化
	fmt.Printf("&newA: %p\n", &newA)  // 但&newA不变

	c := 10
	modInt(c)
	fmt.Printf("&c after mod: %p\n", &c)

	//reflect.Ptr
	//var b *int
	//c := 10
	//b = &c
	//fmt.Printf("&c: %p\n", &c)
	//c++
	//fmt.Printf("&c: %p\n", &c)
	//
	//modInt(c)
	//fmt.Printf("&c after mod: %p\n", &c)
	//
	//fmt.Printf("b: %p\n", b)
	//fmt.Printf("&b: %p\n", &b)
	//modPtr(b)
	//fmt.Printf("b out func after mod: %p\n", b)
	//fmt.Printf("&b out func after mod: %p\n", &b)
	//fmt.Println("b value: ", *b)  // 感觉这个情况和map很像
	////d := map[int]string{1: "aa"}
	////pd := &d
	//
	//
	//
	//fmt.Println("a value: ", a)
	//fmt.Printf("a after mod: %p\n", a)
	//fmt.Printf("&a after mod: %p\n", &a)

}

func modMap(p map[int]string) {
	p[1] = "aa"
	//p = nil
	fmt.Println("p value: ", p)
	fmt.Printf("a in func: %p\n", p)
	fmt.Printf("&a in func: %p\n", &p) // 这里&a变了, 但是a还是没变
}

func modInt(n int) {
	n++
	fmt.Printf("&c in func: %p\n", &n) // 这里的&n 和外边的 &c不一样, 所以我还是觉得&c和&a一样, 确实是存的变量名的地址?
	// 由于传值导致了变量拷贝, 所以地址是不同的
}

func modPtr(a *map[int]string){
	*a = nil
	fmt.Printf("newA in ptrfunc: %p\n", a)  // 由于是指针，对应函数外&newA地址，不变
	fmt.Printf("&newA in ptrfunc: %p\n", &a)  // 对应函数内变量名所在地址
}
