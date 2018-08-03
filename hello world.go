package main

import "fmt"

func main() {
	/* 这是我的第一个简单的程序 */
	//var a string = "123.000"
	a, b := 123, 321
	a, b = b, a
	c := "aaa"
	d := true
	e := &a
	arr1 := [3]int{1, 2, 3}
	//_ , b = a, 222
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("Hello, World!")
	fmt.Printf("%d\n", a)
	fmt.Printf("a的类型为:%T\n", a) // => a的类型为:int
	fmt.Printf("a的单引号值为:%q\n", a) // => a的单引号值为:'Ł'
	fmt.Printf("c的类型为:%T\n", c) // => c的类型为:string
	fmt.Printf("c的值为:%s\n", c) // => c的值为:aaa
	fmt.Printf("c的double-quoted为:%q\n", c) // => c的double-quoted为:"aaa"
	fmt.Printf("d的类型为:%T\n", d) // => d的类型为:bool
	fmt.Printf("e的类型为:%T\n", e) // => e的类型为:*int
	fmt.Printf("e的值为:%d\n", *e) // => e的值为:321
	fmt.Printf("e的#v值为:%#v\n", *e) // => e的#v值为:321
	fmt.Printf("e的+v值为:%+v\n", *e) // => e的+v值为:321
	a = 456
	fmt.Printf("e的新值为:%d\n", *e) // => e的新值为:456
	arr2 := [4]int{}
	for k, v := range arr1 {
		fmt.Printf("arr1 的第 %d 位是 %d\n", k, v) // => arr1 的第 [0,1,2] 位是 [1,2,3]
		arr2[k] = v
	}
	fmt.Printf("arr2 是 %v\n", arr2) // => arr2 是 [1 2 3 0]
	// 以上代码等价于ruby中的
	// arr1.each do |k, v|
	// 	 puts "arr1 的第 #{k} 位是 #{v}"
	// end

	for i := 0; i < 3; i++ {
		fmt.Println("for like c:", arr1[i]) // => for like c: [1-3]
	}
	for a < 460 {
		fmt.Println("for like while in c:", a) // => for like while in c: [456-459]
		a++
	}
	for {
		a++
		fmt.Println("for like c:(;;)", a) // => for like c:(;;) [461-471]
		if a > 470 {
			break
		}
	}
	f := 666
	g := &f
	fmt.Println("f is", f) // => f is 666
	fmt.Println("g is", g) // => g is 0xc42001c1a0
	f = 888
	fmt.Println("g new pos is", g) // => g new pos is 0xc42001c1a0
	fmt.Println("g new value is", *g) // => g new value is 888

	var arr3 = [3]int{1,2,3}
	fmt.Printf("arr3 v# value is %#v\n", arr3) // => arr3 v# value is [3]int{1, 2, 3}
	fmt.Printf("arr3 v+ value is %+v\n", arr3) // => arr3 v+ value is [1 2 3]
}
