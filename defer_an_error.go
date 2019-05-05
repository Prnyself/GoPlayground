package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := ErrorReturn(); err != nil {
		fmt.Println("error in main:", err.Error())
		return
	}
	fmt.Println("no error in main")
}

func ErrorReturn() (err error) {
	defer func() {
		if err != nil {
			fmt.Println("defer:", err.Error())
		} else {
			fmt.Println("no error")
		}
	}()

	a, err := GotError()
	if err != nil {
		fmt.Println("error handle:", err.Error())
		return err
	}
	fmt.Println("a:", a)
	return nil
}

func GotError() (int, error) {
	//return 1, nil
	return 1, errors.New("you got an error")
}
