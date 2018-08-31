// go两种执行方式的区别？
// 1.如果我们先编译生成了可执行文件，这个文件可以处处运行。
// 2. go run 方式需要安装go 的 sdk
// 3. 编译的可执行文件会大一些

// 编译指定输出名称为output
// go build -o output

// Go编程的注意点
// 1.文件以go扩展名结尾 2.应用程序的执行入口是main()函数 3.Go语言严格区分大小写 4.Go每条语句没有分号
// 5.Go是逐行编译的 一行只能写一条语句 6.定义的变量或者import没有用到，Go编译不会通过。7.大括号成对出现 缺一不可



package main

import (
	"fmt"
)

func main () {
	fmt.Println("hello world")
}
