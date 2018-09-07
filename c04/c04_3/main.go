package main

import (
	"fmt"
)

type MethodUtils struct {

}

func (m *MethodUtils) Print() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (m *MethodUtils) Print2(n1 int, n2 int) {
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (m *MethodUtils) Area(len int, width int) int {
	return len * width
}

func (m *MethodUtils) JudgeNum(num int) {
	if num % 2 == 0 {
		fmt.Println("是偶数。。")
		return
	}

	fmt.Println("是奇数。。")
}

func (m *MethodUtils) Transpose(arr [3][3]int) ([3][3]int) {
	var retArr [3][3]int
	for index1, value1 := range(arr) {
		for index2, value2 := range(value1) {
			retArr[index2][index1] = value2
		}
	}

	return retArr
}

type Person struct {
	Name string
}

func (p *Person) test01() {
	p.Name = "jack"
	fmt.Println("test01()=", p.Name)
}

func (p Person) test02() {
	p.Name = "mary"
	fmt.Println("test02()=", p.Name)
}

//方法练习
func main() {
	//1.编写结构体(MethodUtils)，编程一个方法，方法不需要参数，在方法中打印一个 10*8 的矩形，在 main 方法中调用该方法。
	var m MethodUtils
	m.Print()
	//2.编写一个方法，提供 m 和 n 两个参数，方法中打印一个 m*n 的矩形
	m.Print2(5, 5)
	//3.编写一个方法算该矩形的面积(可以接收长 len，和宽 width)，
	//将其作为方法返回值。在 main 方法中调用该方法，接收返回的面积值并打印。
	area := m.Area(5, 5)
	fmt.Println("area=", area)
	//4.编写方法:判断一个数是奇数还是偶数
	m.JudgeNum(10)
	//5.编写一个方法，使给定的二维数组(3*3)转置

	var arr [3][3]int = [3][3]int{{1,2,3}, {4,5,6}, {7,8,9}}
	transArr := m.Transpose(arr)


	fmt.Println(transArr)

	var p Person
	p.test01()//jack
	(&p).test01()//jack 从形式上是传入地址 但仍然是值拷贝
	p.test02()//mary 等价(&p).test02() 形式是传入值 但任然是地址拷贝
	(&p).test02()//mary

	//不管调用形式如何，真正决定是值拷贝还是地址拷贝，看这个方法是和哪个类型绑定.
	//如果是和值类型，比如 (p Person) , 则是值拷贝，如果和指针类型，比如是 (p *Person) 则是地址拷贝。

	//创建结构体变量时指定字段值
	//Golang 在创建结构体实例(变量)时，可以直接指定字段的值
	var p1 = Person{"xiao"}
	p2 := Person{"hello"}

	//下面的写法不依赖字段定义的顺序
	var p3 = Person{
		Name: "jack",
	}

	p4 := Person{
		Name:"xiao",
	}

	fmt.Println(p1, p2, p3, p4)

	//方式二
	//返回结构体的指针类型
	var p5 *Person = &Person{"xiao"}
	p6 := &Person{"first"}

	var p7 = &Person{
		Name: "jack",
	}

	p8 := &Person{
		Name:"xiao",
	}

	fmt.Println(p5, p6, p7, p8)
}

//方法和函数区别
//1.调用方式不一样
//函数的调用方式: 函数名(实参列表)
//方法的调用方式: 变量.方法名(实参列表)
//2.对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
//3.对于方法(如 struct 的方法)，接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以

