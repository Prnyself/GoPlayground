package main

import (
	"fmt"
	"reflect"
)

type order1 struct {
	ordId      int
	customerId int
}

func createQuery1(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%s value:%v\n", i, v.Field(i).Type(), v.Field(i))
		}
	}

}
func main() {
	o := order1{
		ordId:      456,
		customerId: 56,
	}
	createQuery1(o)
}