// go两种执行方式的区别？
// 1.如果我们先编译生成了可执行文件，这个文件可以处处运行。
// 2. go run 方式需要安装go 的 sdk
// 3. 编译的可执行文件会大一些

// go build -o output

package main

import (
	"fmt"
)

func main () {
	fmt.Println("hello world")
}
