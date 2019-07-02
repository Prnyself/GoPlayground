package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func main() {

	var str1, str2 string
	str1 = getCurrentDirectory()

	str2 = getParentDirectory(str1)
	fmt.Println(str2)

	std1 := os.Stdout
	std2 := os.Stdout
	fmt.Printf("std1 == stdout? %v; std2 == stdout? %v; std1 == std2? %v\n",
		std1 == os.Stdout, std2 == os.Stdout, std1 == std2)
}
