package main

import (
	"strconv"
	"fmt"
)

//基本数据类型的相互转换
//Golang 和 java / c 不同，Go 在不同类型的变量之间赋值时需要显式转换。也就是说Golang中数据类型不能自动转换。
//表达式 T(v) 将值 v 转换为类型 T
// T: 就是数据类型，比如 int32，int64，float32 等等 v: 就是需要转换的变量

//基本数据类型相互转换的注意事项
//1. Go中，数据类型的转换可以是从表示范围小-->表示范围大，也可以范围大--->范围小
//2. 被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化!
//3. 在转换中，比如将 int64 转成 int8 【-128---127】 ，编译时不会报错，只是转换的结果是按 溢出处理，和我们希望的结果不一样。 因此在转换时，需要考虑范围.

func main() {
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)

	fmt.Printf("i=%v n1=%v n2=%v n3=%v\n", i, n1, n2, n3)

	fmt.Printf("i type is %T\n", i)

	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=", num2)

	//基本数据类型和 string 的转换
	//1. 基本类型转 string 类型 fmt.Sprintf("%参数", 表达式)
	str := fmt.Sprintf("%d", i)
	fmt.Printf("str type %T str %q\n", str, str)

	str2 := fmt.Sprintf("%f", n1)
	fmt.Printf("str type %T str %q\n", str2, str2)

	isTrue := false
	str3 := fmt.Sprintf("%t", isTrue)
	fmt.Printf("str type %T str %q\n", str3, str3)

	myChar := 'b' 
	str4 := fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str %q\n", str4, str4)

	//使用strconv包的函数
	num3 := 99
	num4 := 23.456
	b2 := true

	str5 := strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str %q\n", str5, str5)

	//f格式 10 表示小数保留10位，64表示这个小数时float64
	str6 := strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str %q\n", str6, str6)

	str7 := strconv.FormatBool(b2)
	fmt.Printf("str type %T str %q\n", str7, str7)

	var num5 int64 = 4567
	str8 := strconv.Itoa(int(num5))
	fmt.Printf("str type %T str %q\n", str8, str8)

	stringToOther()
}

func stringToOther() {
	//2. string 类型转基本类型
	var str string = "true"
	var b bool

	//ParseBool会返回两个值(value bool, err error)
	//我们只想获取value 不想获取err所以用_忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	var str2 string = "123456789"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)

	var str3 string = "123456.2345"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)

	//string 转基本数据类型的注意事项
	//在将 String 类型转成 基本数据类型时，要确保 String 类型能够转成有效的数据，
	//比如 我们可以 把 "123" , 转成一个整数，但是不能把 "hello" 转成一个整数，如果这样做，Golang 直接将其转成 0 ， 
	//其它类型也是一样的道理. float => 0 bool => false
}


