package main

import (
	"fmt"
	"math"
)

var sss = "123"

const (
	a = "123"
	b = len(a)
	c
	d
)

const (
	aa, bb = 1, "222"
	cc, dd
)

const (
	aaa = 'A'
	bbb
	ccc = iota
	ddd
)

const (
	eee = iota
)

const (
	_          = iota
	KB float64 = 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("a=" + a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("aa=", aa)
	fmt.Println("bb=", bb)
	fmt.Println("cc=", cc)
	fmt.Println("dd=", dd)
	fmt.Println("aaa=", aaa)
	fmt.Println("bbb=", bbb)
	fmt.Println("ccc=", ccc)
	fmt.Println("ddd=", ddd)
	fmt.Println("eee=", eee)
	fmt.Println(^-3) //如果是无符号数，则全1进行异或; 有符号数用-1进行异或
	fmt.Println(!true)
	var fff uint8 = 5
	fmt.Println("fff, ", ^fff)
	//var gg int
	//gg = 'a'
	//fmt.Println("gg, ", gg)
	rea, fraction := math.Modf(1.234)
	fmt.Println("real:", rea, ",fraction:", fraction)
	fmt.Println(1 << 10 << 10 >> 10)

	/*
		6: 0110
		11: 1011
		& 0010 = 2
		| 1111 = 15
		^ 1101 = 13
		&^ 0100 = 4  // &^将指定位置为0
	*/
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)

	x := 0
	if x > 0 && (10/x) > 1 {
		fmt.Println("OK")
	}

	fmt.Println("KB=", KB)
	fmt.Println("MB=", MB)
	fmt.Println("GB=", GB)
	fmt.Println("TB=", TB)
	fmt.Println("PB=", PB)
	fmt.Println("EB=", EB)
	fmt.Println("ZB=", ZB)
	fmt.Println("YB=", YB)
	fmt.Println("BB=" + "123")

	//a1, a2 := 10, 0
	//a3 := a1 / a2
	//fmt.Println(a3)

	//k := 6
	//switch k {
	//case 4:
	//	fmt.Println("was <= 4"); fallthrough;
	//case 5:
	//	fmt.Println("was <= 5"); fallthrough;
	//case 6:
	//	fmt.Println("was <= 6"); fallthrough;
	//case 7:
	//	fmt.Println("was <= 7"); fallthrough;
	//case 8:
	//	fmt.Println("was <= 8"); fallthrough;
	//default:
	//	fmt.Println("default case")
	//}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s + "a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}

//
