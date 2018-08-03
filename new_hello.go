package main

import (
	"fmt"
)

func main() {
	s := "hello, world"
	fmt.Println(s)
	fmt.Printf("length of s is %d\n", len(s))
	hello := s[:5]
	world := s[7:]
	fmt.Printf("first: %s, second: %s", hello, world)
}
