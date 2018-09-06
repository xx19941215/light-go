package main

import (
	"fmt"
	"errors"
)

//错误处理
//Go中引入的处理方式为:defer,panic,recover
//这几个异常的使用场景可以这么简单描述:Go 中可以抛出一个 panic 的异常，然后在 defer 中
//通过 recover 捕获这个异常，然后正常处理
func main() {
	// test()
	test03()
	fmt.Println("继续执行")

}

func test() {
	defer func() {
		err := recover() //内置的recover可以捕获异常
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num := 0
	res := num1 / num
	fmt.Println("res=", res)
}
//自定义错误的介绍
//Go 程序中，也支持自定义错误， 使用 errors.New 和 panic 内置函数
//errors.New("错误说明") , 会返回一个 error 类型的值，表示一个错误
//panic 内置函数 ,接收一个 interface{}类型的值(也就是任何值了)作为参数。可以接收error类型的变量，输出错误信息，并退出程序.


func test02() (err error) {
	return errors.New("发生错误")
}

func test03() {
	err := test02()
	if err != nil {
		panic(err)
	}
}