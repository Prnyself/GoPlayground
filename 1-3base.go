// 当前程序包名
package main

// 导入其它包
import "fmt"
import "math"

// 全局变量声明不能省略var关键字
var PI = 3.14

// 全局变量声明与复制
var name = "gopher"

// 一般类型声明
type newType int

// 结构类型声明
type gopher struct{}

// 接口类型声明
type golang interface{}

type byte int8
type rune int32
type 文本 string

// main函数声明
func main() {
	var a [1]bool
	var b 文本
	b = "中文类型名"
	var c int = 1
	d := false
	var x, y, z = 1, 2, 3
	xx, _, zz := 3, 2, 1
	var aa float32 = 100.1
	fmt.Println("Hello Go!")
	bb := int(aa)
	// cc := bool(aa)
	fmt.Println("bb=",bb)

	fmt.Println("a=",a)
	fmt.Println("math.MinInt8=",math.MinInt8)
	fmt.Println("math.MaxInt8=",math.MaxInt8)
	fmt.Println("b=",b)
	fmt.Println("c=",c)
	fmt.Println("d=",d)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(xx)
	fmt.Println(zz)
	fmt.Println("test")
	fmt.Printf("1\r")
	fmt.Printf("2\r")
	fmt.Printf("3\r")
}