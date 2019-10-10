package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "abc123456"
	bPassword := ([]byte)(password)
	//wrongPassword := []byte("123")
	digest, err := bcrypt.GenerateFromPassword(bPassword, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	//err = bcrypt.CompareHashAndPassword(digest, wrongPassword)
	err = bcrypt.CompareHashAndPassword(digest, bPassword)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(digest), password, "compared")
}
