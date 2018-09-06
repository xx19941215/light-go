package main

import (
	"fmt"
)

var age = test()

func test() int {
	fmt.Println("test call")
	return 90
}
//init函数
//每一个源文件都可以包含一个 init 函数，该函数会在 main 函数执行前，
//被 Go 运行框架调用，也 就是说 init 会在 main 函数前被调用。
func init() {
	fmt.Println("init call")
}

func main() {
	fmt.Println("so young", age)
}

//细节说明
//1.如果一个文件同时包含全局变量定义，init 函数和 main 函数，则执行的流程全局变量定义->init函数->main 函数
//2. init函数最主要的作用，就是完成一些初始化的工作
//3. 面试题:案例如果 main.go 和 utils.go 都含有 变量定义，init 函数时，执行的流程 又是怎么样的呢?