package main

import (
	"bufio"
	"fmt"
	"os"
)

// 标准库之flag库--获取命令行参数
func main() {
	fmt.Println("Hello world")

	f := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入字符串: 如需退出请输入exit")
		input, err := f.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		input = input[:len(input)-1]
		fmt.Println("您输入的是:", input)

		if input == "exit" {
			fmt.Println("bye!")
			break
		}
	}
}
