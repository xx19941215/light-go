package main

import (
	"fmt"
)

//数组
//在go中 数组是值类型
//var 数组名 [数组大小]数据类型
func main() {
	var hens [7]float64
	hens[0] = 3.0
	hens[1] = 3.0
	hens[2] = 3.0
	hens[3] = 3.0
	hens[4] = 3.0
	hens[5] = 3.0
	hens[6] = 3.0

	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}

	avgWeight := fmt.Sprintf("%.2f", totalWeight / float64(len(hens)))
	fmt.Printf("totalWeight=%v, avgWeight=%v", totalWeight, avgWeight)

	var arr [3] int
	fmt.Println(arr)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3 
	fmt.Println(arr)
	fmt.Printf("arr的地址是=%p arr[0]的地址%p arr[1]的地址%p arr[2]的地址%p\n", &arr, &arr[0], &arr[1], &arr[2])

	// var score [5]float64
	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("请输入第%v个元素的值\n", i + 1)
	// 	fmt.Scanln(&score[i])
	// }

	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("第%d个元素的值%v", i + 1, score[i])
	// }

	//四种初始化数组的方式
	var numArr01 [3]int = [3]int{1,2,3}
	fmt.Println("numArr01", numArr01)

	var numArr02 = [3]int{5,6,7}
	fmt.Println("numArr02", numArr02)

	var numArr03 = [...]int{8,9,10}
	fmt.Println("numArr03", numArr03)

	var numArr04 = [...]int{1:800, 0:900, 2:999}
	fmt.Println("numArr04", numArr04)

	strArr := [...]string{1:"tom", 0:"jack", 2:"xiaox"}

	fmt.Println("strArr", strArr)

	//数组遍历
	//常规遍历
	//for range遍历
	//这是 Go 语言一种独有的结构，可以用来遍历访问数组的元素。
	//for--range的基本语法
	//1. 第一个返回值index是下标
	//2. 第二个返回值value是该下标位置的值
	//3. 他们都是for循环内部可见的变量
	//4. 如果不想使用index可以使用_代替
	//5. index和value可以取其他名字

	for _, val := range(strArr) {
		fmt.Printf("值分别是%v", val)
	}

	//数组注意事项
	//1.数组是多个相同类型数据的组合,一个数组一旦声明/定义了,其长度是固定的, 不能动态变化
	//2. var arr []int 这时 arr 就是一个 slice 切片，切片后面专门讲解，不急哈.
	//3. 数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不能混用。
	//4. 数组创建后，如果没有赋值，有默认值(零值)
	//5. 使用数组的步骤 1. 声明数组并开辟空间 2 给数组各个元素赋值(默认零值) 3 使用数组
	//6. 数组的下标是从 0 开始的
	//7. 数组下标必须在指定范围内使用，否则报 panic:数组越界，比如
	//8. Go的数组属值类型，在默认情况下是值传递，因此会进行值拷贝。数组间不会相互影响
	//9.长度是数组类型的一部分，在传递函数参数时 需要考虑数组的长度

	var myArr [3]int = [3]int{1,2,3}
	test01(myArr)
	fmt.Println("myArr", myArr)
	test02(&myArr)
	fmt.Println("myArr", myArr)

}

func test01(arr [3]int) {
	arr[0] = 9
	fmt.Println("arr", arr)
}

func test02(arr *[3]int) {
	arr[0] = 9
}

//细节说明
//数组的地址可以通过数组名来获取 &arr
//数组的第一个元素的地址，就是数组的首地址
//数组的各个元素的地址间隔是依据数组的类型决定，比如 int64 -> 8 int32->4...
