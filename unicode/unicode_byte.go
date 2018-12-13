package main

import "fmt"

func main() {
	str1 := "中文"
	fmt.Printf(" => bytes(hex): [% x]\n", []byte(str1)) // 百分号 空格 x 可以将打印的字符串或切片在字节间用空格隔开


	str2 := '0'
	fmt.Printf("%x\n", str2)

	s1 := fmt.Sprintf("%x", str1)
	fmt.Println(s1)
	fmt.Println(len(s1))

	int1 := 1
	fmt.Printf("% d\n", int1)

	int2 := -1
	fmt.Printf("% d\n", int2)
}
