package main

import (
	"fmt"
)

//基本数据类型的默认值
// 整型 0
// 浮点型 0
// 字符型 “”
// 布尔类型 false

func main() {
	var a int
	var b float32
	var c string
	var isTrue bool
	fmt.Printf("a=%d, b=%v, c=%v, isTrue=%v", a, b, c, isTrue)
}