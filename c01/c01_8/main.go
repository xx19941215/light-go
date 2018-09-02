package main

import (
	"fmt"
)

//string类型
//字符串就是一串固定长度的字符连接起来的字符序列。
//Go 的字符串是由单个字节连接起来的。
//Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本
func main() {
	var address string = "hello"	
	fmt.Println(address)
	fmt.Println(address[0]) //打印h的utf-8编码值
	//注意事项
	//1. Go语言的字符串的字节使用UTF-8编码标识Unicode文本，
	//这样Golang统一使用UTF-8编码,中文 乱码问题不会再困扰程序员。
	//2. 字符串一旦赋值了，字符串就不能修改了:在 Go 中字符串是不可变的。address[0] = 's' 会报错
	//3. 字符串的两种表示形式
	// 3.1 双引号, 会识别转义字符
	// 3.2 反引号，以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果

	str3 := `
		func main() {
			fmt.Println(address)
		}
	`
	fmt.Println(str3)
	//4. 字符串拼接方式
	str4 := "xiao"
	str4 += "xiao"

	fmt.Println(str4)

	//5. 当一行字符串太长时，需要使用到多行字符串，可以如下处理
	str5 := "adadas" + "adsadasdd" +
		"sdfsfsdf" + "dfdfsdf" +
		"dsfsfsdf"
	
	fmt.Println(str5)


}