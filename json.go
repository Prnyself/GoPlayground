package main

import (
	"encoding/json"
	"fmt"
)

type JsonTest struct {
	A string `json:"a"`
}

func main() {
	res, err := json.Marshal(map[string]string{
		"a": "a1",
		"b": "b2",
	})
	if err != nil {
		panic(err)
	}

	a := &JsonTest{}
	err = json.Unmarshal(res, a)
	if err != nil {
		panic(err)
	}

	fmt.Printf("a: %+v", a)
}
