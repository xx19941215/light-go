package main

import (
	"fmt"
)

//类型断言
//类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言

type Student struct {
	Name string
}

func main() {
	var x interface{}
	var b2 float32 = 1.1
	x = b2 //空接口 可以接受任意值

	//y := x.(float32)
	//fmt.Printf("y的类型是%T 值是%v", y, y)

	//在进行类型断言时，如果类型不匹配，就会报panic, 因此进行类型断言时，要确保原来的空接口指向的就是断言的类型
	//如何在进行断言时，带上检测机制，如果成功就 ok,否则也不要报 panic

	if y, ok := x.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("y的类型是%T 值是%v", y, y) } else {
		fmt.Println("convert fail")
	}

	fmt.Println("继续执行")

	var n1 float32 = 1.1
	var n2 float64 = 2.2 
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300

	s1 := Student{"xiao"}
	s2 := &Student{"tom"}

	TypeJudge(n1, n2, n3, name, address, n4, s1, s2)
}

//类型断言最佳实践2
//写一函数，循环判断传入参数的类型
func TypeJudge(items...interface{}) {
	for index, v := range(items) {
		switch v.(type) {
			case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index, v)
			case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n", index, v)
			case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n", index, v)
			case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型，值是%v\n", index, v)
			case string:
			fmt.Printf("第%v个参数是字符串类型，值是%v\n", index, v)
			case Student:
			fmt.Printf("第%v个参数是Student类型，值是%v\n", index, v)
			case *Student:
			fmt.Printf("第%v个参数是*Student类型，值是%v\n", index, v)
			default:
			fmt.Printf("第%v个参数不确定类型，值是%v\n", index, v)
		}
	}
}
