package main

import (
	"fmt"
)

func main() {
	var nums []int = []int{2,7,11,15}
	var target int = 9

	ret := twoSum(nums, target)
	fmt.Printf("ret=%v", ret)
}

func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                ret[0] = i
                ret[1] = j
                break
            }
        }
    }
    
    return ret
}