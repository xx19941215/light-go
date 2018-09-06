package main

import (
	"fmt"
)

//
func main() {
	arr := [6]int{1,8, 10, 89, 1000, 1234}
	//测试一把
	binarySearch(&arr, 0, len(arr) - 1, 1)	
}

//
func binarySearch(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("没找到")
		return
	}

	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		binarySearch(arr, leftIndex, middle - 1, findVal)
	} else if (*arr)[middle] < findVal {
		binarySearch(arr, middle + 1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了%v, 下标为%v", findVal, middle)
	}
}