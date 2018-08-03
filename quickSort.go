package main

import "fmt"

func main() {
	res := []int{2, 1, 3, 4, 6, 3, 8, 7}
	sort(res)
	fmt.Println(res)
}

func sort(arr []int) () {
	fmt.Println(arr)
	length := len(arr)
	if length <= 1 {
		return
	}
	flag := 0
	for i := flag + 1; i < length; i++ {
		if arr[i] < arr[flag] {
			tmp := arr[i]
			copy(arr[flag+1:i+1], arr[flag:i])
			arr[flag] = tmp
			flag++
		}
	}
	sort(arr[0:flag])
	sort(arr[flag+1:])
}
