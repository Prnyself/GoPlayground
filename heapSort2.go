package main

import (
	"fmt"
	"container/heap"
)

type myHeapArr []int

func main() {
	arr := &myHeapArr{0, 4, 6, 3, 1, 2, 5, 7, 8}
	fmt.Println(arr, "ori")
	heap.Init(arr)
	heap.Push(arr, 4)
	//old := *arr
	//old[5] = 0
	heap.Fix(arr, 0)
	fmt.Println(arr, "arrayed")
	for arr.Len() > 0 {
		x := heap.Pop(arr)
		fmt.Printf("%d ", x)
	}

	//h := &myHeapArr{2, 1, 5, 100, 3, 6, 4, 5}
	//heap.Init(h)
	//heap.Push(h, 3)
	//heap.Fix(h, 3)
	//fmt.Printf("minimum: %d\n", (*h)[0])
	//for h.Len() > 0 {
	//	fmt.Printf("%d ", heap.Pop(h))
	//}
}

func (h myHeapArr) Len() int           { return len(h) }
func (h myHeapArr) Less(i, j int) bool { return h[i] < h[j] }
func (h myHeapArr) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *myHeapArr) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		fmt.Println("type is not int")
		return
	}
	*h = append(*h, value)
}

func (h *myHeapArr) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
