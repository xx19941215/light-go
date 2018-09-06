package main

import (
	"fmt"
	u "light-go/util"
)

//函数
func main() {
	res := u.Cal(10, 20, '-')
	fmt.Println("res=", res)


	n1 := 10
	n2 := 8
	//使用_忽略值
	_, sub := getSumAndSub(n1, n2)
	fmt.Println("sub=", sub)
	//一共多少桃子
	all := solution(1)
	fmt.Println("all=", all)

	n4 := 10
	n5 := &n4
	test3(n5)
	fmt.Println("n4=", n4)

	type myInt int
	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1)//go认为是两个不同的类型
	fmt.Println("num1=", num1, "num2=", num2)

	type myFuncType func(int, int) int

	var myFuncVar myFuncType = getSum
	res3 := myFun(myFuncVar, 500, 600)

	fmt.Println("sum=", res3);

	res4 := mySum(10, 20, 30, 40)
	fmt.Println("res4=", res4);

	//判读是否可以正常执行,可以
	fmt.Println("sum(1,2)", mySum1(1, 2))
}

//包
//go 的每一个文件都是属于一个包的，也就是说 go 是以包的形式来管理文件和项目目录结构的
//包的作用
//1.区分相同名字的函数、变量等标识符
//2.当程序文件很多时,可以很好的管理项目
//3.控制函数、变量等访问范围，即作用域
//打包 package "包名"
//引包 import "包的路径"
//包细节说明
//1. 在给一个文件打包时，该包对应一个文件夹，比如这里的 util 文件夹对应的包名就是 util, 
//2. 文件的包名通常和文件所在的文件夹名一致，一般为小写字母
//3. 当一个文件要使用其它包函数或变量时，需要先引入对应的包
//4. 在 import 包时，路径从 $GOPATH 的 src 下开始，不用带 src , 编译器会自动从 src 下开始引入
//为了让其它包的文件，可以访问到本包的函数，则该函数名的首字母需要大写，类似其它语言 的 public ,这样才能跨包访问。
//比如 util.go 的Cal
//5. 在访问其它包函数，变量时，其语法是 包名.函数名， 比如这里的 main.go 文件中
//6. 如果包名较长，Go 支持给包取别名， 注意细节:取别名后，原来的包名就不能使用了
//说明: 如果给包取了别名，则需要使用别名来访问该包的函数和变量。
//7. 在同一包下，不能有相同的函数名(也不能有相同的全局变量名)，否则报重复定义
//8. 如果你要编译成一个可执行程序文件，就需要将这个包声明为 main , 即 package main .这个就
// 是一个语法规范，如果你是写一个库 ，包名可以自定义

//关于函数调用的说明
//1.在调用一个函数时，会给该函数分配一个新的空间，编译器会通过自身的处理让这个新的空间 和其它的栈的空间区分开来
//2. 在每个函数对应的栈中，数据空间是独立的，不会混淆
//3. 当一个函数调用完毕之后，程序会销毁这个函数对应的栈空间

//return 返回值
//1. go函数支持多个返回值，这一点是其他编程语言没有的。
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func getSumAndSubNamedReturn(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}
//猴子吃桃问题
//有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个!以后每天猴子都吃其中的一半，然后
//再多吃一个。当到第十天时，想再吃时(还没吃)，发现只有 1 个桃子了。问题:最初共多少个桃子?

func solution(day int) int {
	if day == 10 {
		return 1
	}

	return (solution(day + 1) + 1) * 2;
}

//函数细节再强调
//1. 函数的形参列表可以是多个，返回值列表也可以是多个。
//2. 形参列表和返回值列表的数据类型可以是值类型和引用类型。
//3. 函数的命名遵循标识符命名规范，首字母不能是数字，首字母大写该函数可以被本包文件和其
//它包文件使用，类似 public , 首字母小写，只能被本包文件使用，其它包文件不能使用，类似 privat
//4. 函数中的变量是局部的，函数外不生效
//5. 基本数据类型和数组默认都是值传递的，即进行值拷贝。在函数内修改，不会影响到原来的值。
//6. 如果希望函数内的变量能修改函数外的变量(指的是默认以值传递的方式的数据类型)，可以传 入变量的地址&，函数内以指针的方式操作变量。从效果上看类似引用 。

func test3(n *int) {
	*n = *n + 1
}

//Go函数不支持函数重载
//在 Go 中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量 了。通过该变量可以对函数调用
//函数既然是一种数据类型，因此在 Go 中，函数可以作为形参，并且调用

func myFun(funcvar func(int, int) int, num1 int, num2 int) int {
	return funcvar(num1, num2)
}

//为了简化数据类型定义，Go 支持自定义数据类型
//基本语法:type 自定义数据类型名 数据类型 
//理解: 相当于一个别名 案例:type myInt int 
//这时 myInt 就等价 int 来使用了.

//可变参数
// func sum(args...int) sum int {
// }

//支持1到多个参数
// func sub(n1 int, args...int) sum int {
// }

//说明
//args是slice切片 可以通过args[i]访问各个值
//如果一个函数的形参列表中有可变参数，则可变参数需要放在形参列表最后。

func mySum(n1 int, args...int) (ret int) {
	ret = n1
	for i:=0; i< len(args); i++ {
		ret += args[i]
	}
	return
}

//判断下面的函数是否可以执行
func mySum1(n1, n2 float32) float32 {
	return n1 + n2
}