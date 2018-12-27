package main

import (
	"fmt"
)

func main() {
	var nums []int = []int{2, 7, 11, 15}
	var target int = 9

	var s []string = []string{
		"https://www.aaa.com",
		"https://www.bbb.com",
		"https://www.ccc.com",
		"https://www.ddd.com/",
	}
	ret := twoSum(nums, target)
	ret2 := contains(s, "https://www.ddd.com/")
	fmt.Printf("ret=%v", ret)
	fmt.Printf("ret=%v", ret2)
}

func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ret[0] = i
				ret[1] = j
				break
			}
		}
	}

	return ret
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
