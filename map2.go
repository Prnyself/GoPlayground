package main

import "fmt"

type HashMap map[string]bool

func main() {

	m0 := HashMap{}
	m0["0"] = false
	fmt.Println(m0)

	m1 := &HashMap{}
	(*m1)["1"] = false
	fmt.Println(m1)

	var m2 *HashMap
	//m2 := new(HashMap)
	//*m2 = m0   //这样是错误的，因为m2还没有初始化，所以对m2的值的操作会报错
	m2 = &m0
	(*m2)["2"] = false
	fmt.Println(m2)


}
