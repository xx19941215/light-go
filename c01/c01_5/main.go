// 小数类型/浮点型
package main

import (
	"fmt"
)
func main () {
	var price float32 = 89.12
	fmt.Println("price=", price)
	//小数类型分类
	// 单精度float32 占用4字节 表示范围 -3.403E38 - 3.403E38
	// 双精度float64 占用8字节 表示范围 -1.798E308 - 1.798E308
	// 说明
	// 1.浮点数在机器中存放形式的简单说明 浮点数=符号位+指位数+尾数位(浮点数都是有符号的)
	// 2.位数部分可能丢失，造成精度损失
	// 3.float64精度高于float32

	var num1 float32 = -1230.000089
	var num2 float64 = -1230.000089

	fmt.Println("num1=", num1, "num2=", num2)

	//浮点数使用细节
	//1. Go的浮点数类型有固定的范围和长度字段，不受OS的影响
	//2. Go的浮点默认类型声明为float64
	//3. 浮点型常量有两种表示方式
	// 3.1 十进制表示
	var num3 = .52
	var num4 = 5.52
	fmt.Println("num3=", num3, "num4=", num4)
	// 3.2 科学计数法表示
	num5 := 5.123e2 
	num6 := 5.123E2
	num7 := 5.123E-2 //除以10的2次方
	fmt.Println("num5=", num5, "num6=", num6, "num7=", num7)
	//4. 通常情况下 我们应该使用float64
}