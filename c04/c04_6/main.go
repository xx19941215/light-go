package main

import (
	"fmt"
)

//接口
//按顺序,我们应该讲解多态,但是在讲解多态前,我们需要讲解接口(interface)，
//因为在 Golang 中 多态特性主要是通过接口来体现的。
//interface 类型可以定义一组方法，但是这些不需要实现。并且 interface 不能包含任何变量。到某个
//自定义类型(比如结构体 Phone)要使用的时候,在根据具体情况把这些方法写出来(实现)。
//type 接口名 interface {
//method1(参数列表)返回值列表
//method2(参数列表)返回值列表
//}
//func (t 自定义类型) method1(参数列表)(返回值列表) {
//
//}
//func (t 自定义类型) method2(参数列表)(返回值列表) {
//
//}
type AInterface interface {
	BInterface
	CInterface
	Say()
}

type BInterface interface {
	Hello()
}

type CInterface interface {
	C()
}

type Student struct {
	Name string
}

func (s Student) Say() {
	fmt.Println("Student Say")
}

func (s Student) Hello() {
	fmt.Println("Student Hello")
}


func (s Student) C() {

}

type integer int

func (i integer) Say() {
	fmt.Println("integer say", i)
}

func (i integer) Hello() {
	fmt.Println("integer hello")
}

func (i integer) C() {
	fmt.Println("C")
}

func main() {
	//1.接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的
	//多态和高内聚低偶合的思想。
	//2.Golang中的接口，不需要显式的实现。只要一个变量，含有接口类型中的所有方法，那么这个
	//变量就实现这个接口。因此，Golang 中没有 implement 这样的关键字

	//接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量(实例)

	var stu Student
	var a AInterface = stu
	a.Say()

	//接口中所有的方法都没有方法体,即都是没有实现的方法。
	//在 Golang 中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现 了该接口。
	//一个自定义类型只有实现了某个接口，才能将该自定义类型的实例(变量)赋给接口类型
	//只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型。

	var i integer = 10
	var b AInterface = i
	b.Say()

	//一个自定义类型可以实现多个接口
	var stu2 Student
	var a2 AInterface = stu2
	var b2 BInterface = stu2

	a2.Say()
	b2.Hello()

	// Golang接口中不能有任何变量
	// 一个接口(比如 A 接口)可以继承多个别的接口(比如 B,C 接口)，
	//这时如果要实现 A 接口，也必 须将 B,C 接口的方法也全部实现

	//interface类型默认是一个指针(引用类型)，如果没有对interface初始化就使用，那么会输出nil
	//空接口interface{} 没有任何方法，所以所有类型都实现了空接口, 即我们可以把任何一个变量 赋给空接口。
	var t T = stu
	fmt.Println(t)

	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	t = num1
	fmt.Println(t2, t)
}

type T interface {}

//细节说明