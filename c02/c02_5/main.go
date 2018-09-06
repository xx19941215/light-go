package main

import (
	"fmt"
)

//defer
//在函数中，程序员经常需要创建资源(比如:数据库连接、文件句柄、锁等) ，为了在函数执行完毕后，及时的释放资源，
//Go 的设计者提供 defer (延时机制)。
func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}

func sum(n1 int, n2 int) int {
	//当执行到defer时暂不执行 会将defer后的语句压入defer栈
	//当函数执行完毕之后，按照先入后出的方式依次执行
	defer fmt.Println("ok n1", n1)
	defer fmt.Println("ok n2", n2)

	//defer值拷贝
	n1++
	n2++

	res := n2 + n1
	fmt.Println("ok res", res)
	return res
}

//细节说明
//1.当 go 执行到一个 defer 时，不会立即执行 defer 后的语句，而是将 defer 后的语句压入到一个栈 中
//[为了理解，暂时称该栈为 defer 栈], 然后继续执行函数下一个语句。
//2.当函数执行完毕后，在从 defer 栈中，依次从栈顶取出语句执行(注:遵守栈 先入后出的机制)
//3.在 defer 将语句放入到栈时，也会将相关的值拷贝同时入栈。

//最佳实践
//1.defer 最主要的价值是在，当函数执行完毕后，可以及时的释放函数创建的资源。
//2.在 golang 编程中的通常做法是，创建资源后，比如(打开了文件，获取了数据库的链接，或者是
//锁资源)， 可以执行 defer file.Close() defer connect.Close()
//3.在 defer 后，可以继续使用创建资源.
//4.当函数完毕后，系统会依次从 defer 栈中，取出语句，关闭资源.
//5.这种机制，非常简洁，程序员不用再为在什么时机关闭资源而烦心
