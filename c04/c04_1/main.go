package main

import (
	"fmt"
)

//结构体

//Golang也支持面向对象编程(OOP)，
//但是和传统的面向对象编程有区别，并不是纯粹的面向对 象语言。所以我们说 Golang 支持面向对象编程特性是比较准确的。
//Golang没有类(class)，Go语言的结构体(struct)和其它编程语言的类(class)有同等的地位，
//你可以理解 Golang 是基于 struct 来实现 OOP 特性的。
//Golang面向对象编程非常简洁，去掉了传统OOP语言的继承、方法重载、构造函数和析构函数、隐藏的this指针等等
//Golang仍然有面向对象编程的继承，封装和多态的特性，只是实现的方式和其它OOP语言不
//一样，比如继承 :Golang 没有 extends 关键字，继承是通过匿名字段来实现。
//Golang面向对象(OOP)很优雅，OOP本身就是语言类型系统(typesystem)的一部分，通过接口 (interface)关联，耦合性低，也非常灵活。
//后面同学们会充分体会到这个特点。也就是说在 Golang 中面
//向接口编程是非常重要的特性。
type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
}

func main() {
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 10
	cat1.Color = "白"
	cat1.Hobby = "吃"

	fmt.Println(cat1)
//创建结构体变量和访问结构体字段
//1.方式 1-直接声明
//案例演示: var person Person 前面我们已经说了。
//方式 2-{}

	p2 := Cat{"hhh", 20, "黑", "吃"}
	fmt.Println(p2)

	//方式 3-&
	var c3 *Cat = new(Cat)
	(*c3).Name = "dfsf"
	(*c3).Age = 10
	c3.Color = "红"
	(*c3).Hobby = "吃"

	fmt.Println(c3)

	//方式 4-{}
	var c4 *Cat = &Cat{}
	(*c4).Name = "dfsf"
	c4.Age = 100
	c4.Color = "红"
	c4.Hobby = "吃"

	fmt.Println(c4)

	//第 3 种和第 4 种方式返回的是 结构体指针。
	//结构体指针访问字段的标准方式应该是:(*结构体指针).字段名 ，比如 (*person).Name = "tom"
	//但 go 做了一个简化，也支持 结构体指针.字段名, 比如 person.Name = "tom"。更加符合程序员
 	//使用的习惯，go 编译器底层 对 person.Name 做了转化 (*person).Name。

	
	
	

}

//细节说明
//结构体和结构体变量(实例)的区别和联系
//1. 结构体是自定义的数据类型，代表一类事物.
//2. 结构体变量(实例)是具体的，实际的，代表一个具体变量

//如何声明结构体
//type 结构体名称 struct {
//field1 type
//field2 type }

//从概念或叫法上看: 结构体字段 = 属性 = field
//字段是结构体的一个组成部分，一般是基本数据类型、数组,也可是引用类型。
//比如我们前面定义猫结构体 的 Name string 就是属性

//注意事项
//1.字段声明语法同变量，示例:字段名 字段类型
//2.字段的类型可以为:基本类型、数组或引用类型
//3.在创建一个结构体变量后，如果没有给字段赋值，都对应一个零值(默认值)，规则同前面一样:
//布尔类型是 false ，数值是 0 ，字符串是 ""。
//数组类型的默认值和它的元素类型相关，比如 score [3]int 则为[0, 0, 0]
//指针，slice，和 map 的零值都是 nil ，即还没有分配空间。
//不同结构体变量的字段是独立，互不影响，一个结构体变量字段的更改，不影响另外一个, 结构体是值类型。

