package main

import (
	"strconv"
	"fmt"
	"time"
)

//时间日期函数
// 时间和日期相关函数，需要导入 time 包
func main() {
	//当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)

	//获取年月日时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//格式化日期时间
	//1.就是使用 Printf 或者 SPrintf
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(),
	int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	//2.使用 time.Format() 方法完成:
	fmt.Printf(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	//这个字符串的各个数字是固定的，必须是这样写。
	//这个字符串各个数字可以自由的组合，这样可以按程序需求来返回时间
	fmt.Printf(now.Format("15|04|05"))

	//时间的常量
	
	i := 0
	for {
		i++
		fmt.Println(i)
		//休眠
		time.Sleep(time.Microsecond * 100)

		if i == 100 {
			break
		}
	}

	//time的Unix和UnixNano的方法
	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", now.Unix(), now.UnixNano())
	//前者表示Unix时间，即从1970年1月1日 UTC时间到目前的经过的秒数
	//后者是表示的纳秒数 如果纳秒位单位时间超过了int64表示的范围，结果就会是未定义的

	start := time.Now().Unix()
	// test03()
	end := time.Now().Unix()
	fmt.Printf("耗费时间%d s\n", end - start)

	num1 := 100
	fmt.Printf("num1的类型%T, num1的值%v, num1的地址%v\n", num1, num1, &num1)

	num2 := new(int)
	*num2 = 200

	fmt.Printf("num2的类型%T, num2的值%v, num2的地址%v, num2这个指针指向的值%v\n", num2, num2, &num2, *num2)
}

//细节说明

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

//内置函数
//Golang 设计者为了编程方便，提供了一些函数，这些函数可以直接使用，我们称为 Go 的内置函 数。
//文档:https://studygolang.com/pkgdoc -> builtin
//1.len:用来求长度，比如string、array、slice、map、channel
//2.new:用来分配内存，主要用来分配值类型，比如int、float32,struct...返回的是指针
//make:用来分配内存，主要用来分配引用类型，比如channel、map、slice。