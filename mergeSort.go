package main

import "fmt"

func main() {
	res := []int{2, 1, 3, 4, 6, 3, 8, 7, 1}
	fmt.Println(mergeSort(res))
}

func mergeSort(arr []int) ([]int) {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	tmp1 := mergeSort(arr[0:length/2])
	tmp2 := mergeSort(arr[length/2:])
	return merge(tmp1, tmp2)
}

func merge(arr1 []int, arr2 []int) []int {
	result := make([]int, 0)
	for {
		if len(arr1) == 0 || len(arr2) == 0 {
			break
		}
		if arr1[0] < arr2[0] {
			result = append(result, arr1[0])
			//copy(arr1, arr1[1:])
			arr1 = arr1[1:]
		}else{
			result = append(result, arr2[0])
			//copy(arr2, arr2[1:])
			arr2 = arr2[1:]
		}
	}

	if len(arr1) > 0 {
		result = append(result, arr1...)
	} else {
		result = append(result, arr2...)
	}

	return result
}
