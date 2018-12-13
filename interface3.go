package main

import "fmt"

type ReadWriter interface {
	Reader
	Writer
}

type Reader interface {
	Read() (n int)
}

type Writer interface {
	Write(m int)
}

type TriNumber struct {
	hundred int
	ten     int
	one     int
}

func (num TriNumber) Read() int {
	return num.hundred*100 + num.ten*10 + num.one
}

func (num *TriNumber) Write(n int) {
	num.one = n / 100
	num.ten = (n / 10) % 10
	num.hundred = n % 10
}

func main() {
	//var number ReadWriter = new(TriNumber)
	//m := 123
	//number.Write(m)
	//fmt.Println(number.Read())
	//
	//var numRead Reader = number
	//fmt.Println(numRead.Read())
	var numRead Reader = &TriNumber{4, 5, 6}
	fmt.Println(numRead.Read())
	fmt.Printf("%T", numRead)

	//var number ReadWriter = numRead

}
