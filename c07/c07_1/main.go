package main

import (
	"fmt"
	"reflect"
)

//反射
//1.反射可以在运行时动态获取变量的各种信息, 比如变量的类型(type)，类别(kind)
//2.如果是结构体变量，还可以获取到结构体本身的信息(包括结构体的字段、方法)
//3.通过反射，可以修改变量的值，可以调用关联的方法。
//4.使用反射，需要 import (“reflect”)
func main() {
	var num int = 100
	reflectTest01(num)

	stu := Student{
		Name: "Tom",
		Age:  20,
	}

	reflectTest02(stu)

	testInt(&num)
	fmt.Println("num=", num)
}

type Student struct {
	Name string
	Age  int
}

func reflectTest01(b interface{}) {
	//通过反射获取传入的type,kind,值

	//1.先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)

	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)

	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	//将rVal转换成interface{}
	iV := rVal.Interface()
	//将interface{}通过断言转换成需要的类型
	num2 := iV.(int)

	fmt.Println("num2=", num2)

}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	rVal := reflect.ValueOf(b)

	// 获取变量对应的Kind
	kind1 := rVal.Kind()
	kind2 := rType.Kind()

	fmt.Printf("kind=%v kind=%v\n", kind1, kind2)

	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iV, iV)

	stu, ok := iV.(Student)

	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	fmt.Printf("val type=%T\n", val)
	val.Elem().SetInt(110)
	fmt.Printf("val=%v\n", val)
}

//细节说明
//1.reflect.Value.Kind，获取变量的类别，返回的是一个常量
//2.Type 和 Kind 的区别
//Type 是类型, Kind 是类别， Type 和 Kind 可能是相同的，也可能是不同的.
//比如: var num int = 10 num的Type是int , Kind也是int
//比如: var stu Student stu 的 Type 是 pkg1.Student , Kind 是 struct
//3.通过反射可以在让变量在interface{}和Reflect.Value之间互相转换。
//4.使用反射的方式获取变量的值，要求数据类型匹配，比如x是int，那么就应该使用reflect.Value(x).Int()
//5.通过反射的来修改变量, 注意当使用 SetXxx 方法来设置需要通过对应的指针类型来完成, 这样
//才能改变传入的变量的值, 同时需要使用到 reflect.Value.Elem()方法
