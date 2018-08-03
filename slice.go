package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		v := 1
		fmt.Println(&v)
	}

	// 使用数组下标切片
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)     // => [1 2 3 4 5 6 7 8 9 0]
	fmt.Println(arr[5:]) // => [6 7 8 9 0]
	fmt.Println(arr[:5]) // => [1 2 3 4 5]

	// 使用make创建切片
	s := make([]int, 3, 10)
	fmt.Println(s, len(s), cap(s))
	//s[7] = 10
	//fmt.Println(s, len(s), cap(s))

	slice1 := []int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}
	fmt.Printf("slice1: %T, slice2: %T\n", slice1, arr2)

	slice1 = append(slice1, 0)
	copy(slice1[3:], slice1[2:])
	fmt.Println(slice1)
	slice1[2] = 5
	fmt.Println(slice1)

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a[1:3])
	a = append(a[:0], a[1:]...) // 删除开头1个元素
	fmt.Println(a)
	a = append(a[:0], a[3:]...) // 删除开头N个元素
	fmt.Println(a)

	a = []int{1, 2, 3, 4, 5, 6}
	n := copy(a, a[1:])
	//a = a[:copy(a, a[1:])] // 删除开头1个元素
	fmt.Println(n)
	fmt.Println(a)
	a = a[:copy(a, a[2:])] // 删除开头N个元素
	fmt.Println(a)

	x := s[:]
	fmt.Println(x, len(x), cap(x))

	xx := [3]int{1, 2, 3}
	fmt.Println("before func xx is: ", xx)
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println("in func xx is: ", arr)
	}(xx)
	fmt.Println("after func xx is: ", xx)

	fmt.Println("before func a is: ", a)
	func(slice []int) {
		slice[0] = 20
		fmt.Println("in func a is: ", slice)
	}(a)
	fmt.Println("after func a is: ", a)

	mapTest := make(map[int]string, 10)

	mapTest = map[int]string{
		3: "gamma",
		4: "comma",
		5: "sigma",
	}

	mapTest[1] = "alpha"
	mapTest[2] = "beta"
	for k, v := range mapTest {
		fmt.Println("map key is: ", k, ", value is: ", v)
	}

	str := "hello"
	substr := str[1:3]
	fmt.Println(substr)
	fmt.Printf("type of substr is %T\n", substr)
}
