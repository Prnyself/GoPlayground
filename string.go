package main

import (
	"fmt"
	"strings"
)

func main() {
	a := `aaa`
	fmt.Printf("%s, %T\n", a, a)
	path := "www.baidu.com?"
	fmt.Printf("length of %s is: %d, and the last ? is at %d\n", path, len(path), strings.LastIndex(path, "?"))

	fmt.Println(path[:2] + "abc")
}
