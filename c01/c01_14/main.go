package main

import (
	"fmt"
)

//关系运算符
//关系运算符的结果都是 bool 型，也就是要么是 true，要么是 false
//关系表达式经常用在if结构的条件中或循环结构的条件中
func main() {
	var flag bool
	flag = 5 > 4
	fmt.Println(flag);
}

//关系运算符细节说明
//关系运算符的结果都是 bool 型，也就是要么是 true，要么是 false。
//关系运算符组成的表达式，我们称为关系表达式: a > b
//比较运算符"=="不能误写成 "=" !!
