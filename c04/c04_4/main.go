package main

import (
	"fmt"
	"light-go/model"
)

//构造函数
//Golang 的结构体没有构造函数，通常可以使用工厂模式来解决这个问题。
//使用工厂模式实现跨包创建结构体实例(变量)

func main() {
	var s = model.NewStudent("xiao")
	fmt.Println(*s)
	fmt.Println("name=", s.Name)
}

//细节说明