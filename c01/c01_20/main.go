package main

import (
	"fmt"
)

//for循环
func main() {
	var count int
	for count = 0; count < 10; count++ {
		fmt.Println("hello world")
	}

	//第二种写法
	//也可以使用 for ; ; {} 是一个无限循环， 通常需要配合 break 语句使用
	for count < 20 {
		fmt.Printf("hello world %v\n", count)
		count++
	}

	//Golang 提供 for-range 的方式，可以方便遍历字符串和数组(注: 数组的遍历，请看后面的章节)
	//传统方式
	var str string = "hello world呵呵"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}
	//for-range
	var str2 string = "hello world"
	for index, val := range str2 {
		fmt.Printf("index=%d value=%c\n", index, val)
	}

	//注意
	//如果我们的字符串含有中文，那么传统的遍历字符串方式(for range是按照字符遍历所以没有问题)，就是错误，会出现乱码。原因是传统的
	//对字符串的遍历是按照字节来遍历，而一个汉字在 utf8 编码是对应 3 个字节。 
	//如何解决 需要要将 str 转成 []rune 切片.

	str3 := []rune(str)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("%c \n", str3[i])
	}

	//go没有while和do while
	//可以使用for取代a
	//九九乘法表案例
	var num1 byte
	for num1 = 1; num1 < 10; num1++ {
		var num2 byte
		for num2 = 1; num2 <= num1; num2++ {
			fmt.Printf("%d * %d = %d\t", num2, num1, num1 * num2)
		}
		fmt.Println()
	}


}

//细节说明
