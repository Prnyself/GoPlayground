package main

import (
	"os"
	"fmt"
	"bufio"
	"runtime"
	"sync"
	"io"
)

func main() {
	//writeFile()
	res := make([][]string, 5)
	for i := 1; i <= 5; i++ {
		fileName := fmt.Sprintf("res%d.txt", i)
		pstr, err := readFile(fileName)
		if err != nil {
			fmt.Println(err.Error())
		}
		res[i-1] = pstr
	}

	for j := 1; j <= 5; j++ {
		fmt.Println("res1与res", j, "相比:")
		if isEqual(res[0], res[j-1]) {
			fmt.Println("equal")
		} else {
			fmt.Println("not equal")
		}
	}
	//fileName3 := "res3.txt"
	//res3, err := readFile(fileName3)
	//if err!=nil{
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(res3)
	//
	//
	//fileName2 := "res2.txt"
	//res2, err := readFile(fileName2)
	//if err!=nil{
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(res2)
	//
	//if isEqual(res2, res3){
	//	fmt.Println("equal")
	//} else {
	//	fmt.Println("not equal")
	//}
}

func writeFile() {
	fileName := "res5.txt"
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败: ", err.Error())
		return
	}

	w := bufio.NewWriter(file) //创建新的 Writer 对象

	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(30000)

	for i := 0; i < 30000; i++ {
		go func(i int) {
			defer wg.Done()
			_, err2 := w.WriteString(fmt.Sprintln(i))
			if err2 != nil {
				fmt.Println(i, "写入失败: ", err2.Error())
			}
		}(i)
	}
	wg.Wait()
	w.Flush()
}

func readFile(name string) ([]string, error) {
	//res := make([]int, 1000)
	buf := make([]string, 0)
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("打开文件失败:", err.Error())
		return buf, err
	}
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, isPrefix, err1 := r.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return buf, err1
		}
		str := string(line) //转换字符数组为字符串
		buf = append(buf, str)
	}
	return buf, nil

}

func isEqual(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}
	res := true
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			res = false
			break
		}
	}
	return res
}
