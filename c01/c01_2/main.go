// Go 语言的转义字符
// \t 一个制表位
// \n 换行
// \\ 一个\
// \" 一个“
// \r 表示从当前行最前面开始输出覆盖掉之前的内容

package main

import "fmt"

func main() {
	fmt.Println("escaped\tchar")
	fmt.Println("hello\nworld")
	fmt.Println("C:\\Users\\Administrator")
	fmt.Println("xiao说\"你好\"")
	fmt.Println("xiao说\r你好")
}