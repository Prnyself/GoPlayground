package main

import "fmt"

type B []byte
type S string

func main() {
	var b = B([]byte("hello"))
	P(b)
	Pr([]byte("hello"))

	var s = S("world")
	SS(string(s))
	SSr(s)
}

func P(b []byte) {
	fmt.Println("Print in P")
	fmt.Println(b)
}
func Pr(b B) {
	fmt.Println("Print in Pr")
	fmt.Println(b)
}

func SS(s string) {
	fmt.Println("Print in SS")
	fmt.Println(s)
}

func SSr(s S) {
	fmt.Println("Print in SSr")
	fmt.Println(s)
}
