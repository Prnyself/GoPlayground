package main

import (
	"fmt"
)

type I interface {
	Set(uu *User)
}
type User struct {
	Name string
}

func (u *User) Set(uu *User) {

	*u = *uu

}

func main() {
	var u = new(User)
	u.Name = "Jack"

	fmt.Println(*u)
	Copy(u)

	fmt.Println(*u)

}

func Copy(i interface{}) {
	var u User
	u.Name = "Mike"

	uu, ok := i.(*User)
	if !ok {
		fmt.Println("the type is wrong!")
	}

	switch i.(type) {
	case *User:
		fmt.Println("type is *User")
	case int:
		fmt.Println("type is int")
	default:
		fmt.Println("unknown type")
	}

	uu.Set(&u)
	i = uu

	fmt.Println(i)
}
