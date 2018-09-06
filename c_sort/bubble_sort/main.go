package main

import (
	"fmt"
)

func main() {
	arr := [5]int{12, 34, 3, 4, 45}
	bubbleSort(&arr)
	fmt.Println(arr)
}

func bubbleSort(arr *[5]int) {
	temp := 0

	for i := 0; i < len(*arr) - 1; i++ {
		for j := 0; j < len(*arr) - 1 - i; j++ {
			if (*arr)[j + 1] < (*arr)[j] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = temp
			}
		}
	}
}
