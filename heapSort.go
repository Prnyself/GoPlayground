package main

import (
	"fmt"
)

type Arr []int

func main() {
	a := Arr{0, 4, 6, 3, 1, 2, 5, 7, 8}
	fmt.Println(a)
	fmt.Println(len(a))

	heapSort(a, len(a))
	//down(a, 0, len(a))
	fmt.Println(a)

}

func leftChild(index int) int {
	return 2*index + 1
}

func down(arr Arr, index int, length int) {
	var child int
	for i := index; leftChild(i) < length; i = child {

		child = leftChild(i)
		if child != length-1 && arr[child+1] > arr[child] {
			child++
		}

		if arr[i] < arr[child] {
			arr.swap(i, child)
		} else {
			break
		}
	}
	//res, ok := arr[index]
}

func (arr Arr) swap(a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func heapSort(arr Arr, length int) {
	//步骤一：创建大根堆
	for i := length / 2; i >= 0; i-- {
		down(arr, i, length);
	}

	//步骤二：调整大根堆
	for i := length - 1; i > 0; i-- {
		//首尾交换
		arr.swap(0, i)
		//元素下沉
		down(arr, 0, i)
	}

}
