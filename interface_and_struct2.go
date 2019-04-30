package main

type Addder interface {
	add(int)
}
type IntStruct struct {
	num int
}

func (i *IntStruct) add(n int) {
	i.num += n
}

type IntAdd struct {
	Addder
	myNum int
}

func (i *IntAdd) add(n int) {
	i.myNum += 10 * n
}

func main() {
	newIntStruct := &IntStruct{num: 1}
	newIntStruct.add(10) // 11
	println(newIntStruct.num)
	newIntAdd := &IntAdd{}
	newIntAdd.Addder = newIntStruct // newIntStruct类型IntStruct实现了Adder接口
	newIntAdd.add(10)
	println(newIntStruct.num) // 11
	println(newIntAdd.myNum)  // 100
}
