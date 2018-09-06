package main

import (
	"strings"
	"fmt"
)

//闭包
//闭包就是一个函数和与其相关的引用环境组合的一个整体(实体)
func main() {
	// f := AddUpper()
	f := AddUpper2()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后", f2("winter"))
	fmt.Println("文件名处理后", f2("bird.jpg"))
}

func AddUpper() func (int) int {
	var n int = 10
	return func (x int) int {
		n = n + x
		return n
	}
}

//细节说明
//1.AddUpper 是一个函数，返回的数据类型是 fun (int) int
//2.返回的是一个匿名函数, 但是这个匿名函数引用到函数外的n,因此这个匿名函数就和n形成一 个整体，构成闭包。
//3.大家可以这样理解: 闭包是类, 函数是操作，n 是字段。函数和它使用到 n 构成闭包
//4.当我们反复的调用 f 函数时，因为 n 是初始化一次，因此每调用一次就进行累计。
//5.我们要搞清楚闭包的关键，就是要分析出返回的函数它使用(引用)到哪些变量，因为函数和它引
//用到的变量共同构成闭包。
//6.对上面代码的一个修改，加深对闭包的理解

func AddUpper2() func (int) int {
	var n int = 20
	var str string = "hello"

	return func (x int) int {
		n = n + x
		str += string(36)
		fmt.Println("str=", str)
		return n
	}
}

//闭包案例
func makeSuffix(suffix string) func (string) string {
	//如果name没有指定后缀，就加上，否则返回原来的名字
	return func (name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}