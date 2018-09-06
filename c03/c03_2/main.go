package main

import (
	"time"
	"fmt"
	"math/rand"
)

//数组应用
func main() {
	//创建一个 byte 类型的 26 个元素的数组，分别 放置'A'-'Z‘。使用 for 循环访问所有元素并打印
	//出来。提示:字符数据运算 'A'+1 -> 'B'

	var ele [26]byte
	for ind, _ := range(ele) {
		ele[ind] = 'A' + byte(ind)
	}

	for _, val := range(ele) {
		fmt.Printf("%c,", val)
	}

	//请求出一个数组的最大值，并得到对应的下标。
	var intArr [6]int = [...]int{12, 4, 45, 56, 98, 56}
	maxVal := intArr[0]
	maxIndex := 0

	for ind, val := range(intArr) {
		if val > maxVal {
			maxVal = val
			maxIndex = ind
		}
	}

	fmt.Printf("maxVal=%d, maxIndex=%d", maxVal, maxIndex)
	//随机生成五个数，并将其反转打印

	var intArr3 [5]int
	len := len(intArr3)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr3[i] = rand.Intn(100)
	}

	fmt.Println("交换前", intArr3)
	temp := 0
	for i := 0; i < len/2; i++ {
		temp = intArr3[len - 1 - i]
		intArr3[len - 1 - i] = intArr3[i]
		intArr3[i] = temp
	}

	fmt.Println("交换后", intArr3)
}

//细节说明