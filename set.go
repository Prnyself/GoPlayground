package main

import (
	"GoPlayground/mySet"
	"fmt"
)

func main() {
	testSet := mySet.Set{}
	testSet[1] = true
	fmt.Println(testSet[2]) // false

	testSet.Add("a")
	fmt.Println(testSet, testSet.Len(), testSet.IsExist(1)) // map[1:true a:true] 2 true

	testSet.Delete(1)
	fmt.Println(testSet, testSet.Len(), testSet.IsExist(1)) // map[a:true] 1 false
}
