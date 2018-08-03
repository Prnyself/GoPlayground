package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func read0(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	return string(f)
}

func read1(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

func read2(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 0)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf) //n为读了多少个字节
		fmt.Println("每次读了多少个字节:", n)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

func read3(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func main() {

	file := "res1.txt"
	//fmt.Printf("%s", read0(file))
	//fmt.Println("-----------")
	//fmt.Printf("%s", read1(file))
	//fmt.Println("-----------")
	//fmt.Printf("%s", read2(file))
	//fmt.Println("-----------")
	//fmt.Printf("%s", read3(file))
	//fmt.Println("-----------")
	start := time.Now()

	read0(file)
	t0 := time.Now()
	fmt.Printf("Cost time %v\n", t0.Sub(start))

	read1(file)
	t1 := time.Now()
	fmt.Printf("Cost time %v\n", t1.Sub(t0))

	read2(file)
	t2 := time.Now()
	fmt.Printf("Cost time %v\n", t2.Sub(t1))

	read3(file)
	t3 := time.Now()
	fmt.Printf("Cost time %v\n", t3.Sub(t2))

}