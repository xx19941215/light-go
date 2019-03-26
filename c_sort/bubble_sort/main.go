package main

import (
	"fmt"
)

func main() {
	arr := [9]int{12, 34, 3, 4, 45, 100, 34, 88, 76}
	bubbleSort(&arr)
	fmt.Println(arr)
}

func bubbleSort(arr *[9]int) {
	temp := 0
	len := len(arr)

	//外层进行len-1次循环
	for i := 0; i < len-1; i++ {
		swapped := false
		//内循环比较len-i-i次
		for j := 0; j < len-1-i; j++ {
			if arr[j+1] < arr[j] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				swapped = true
			}
		}
		//没有发生交换，说明有序 直接结束算法
		if swapped == false {
			break
		}
	}
}
