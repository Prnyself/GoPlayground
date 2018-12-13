package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func Sha1(data string) string {
	encrypts := sha1.New()
	encrypts.Write([]byte(data)) // Cannot convert expression of type string to type []byte
	return hex.EncodeToString(encrypts.Sum([]byte("")))
}

func main() {
	fmt.Println(Sha1("123456"))
}
