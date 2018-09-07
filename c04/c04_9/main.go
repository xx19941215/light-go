package main

import (
	"fmt"
)

//多态
//变量(实例)具有多种形态。面向对象的第三大特征，
//在 Go 语言，多态特征是通过接口实现的。可 以按照统一的接口来调用不同的实现。这时接口变量就呈现不同的形态。
type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

type Computer struct {

}

//类型断言最佳实践之一
func (c Computer) Working(usb Usb) {
	usb.Start()
	//如果 usb 是指向 Phone 结构体变量，则还需要调用 Call 方法
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}

	usb.Stop()
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

func (p Phone) Call() {
	fmt.Println("手机在打电话")
}



type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

func main() {
	//给 Usb 数组中，存放 Phone 结构体 和 Camera 结构体变量
	var usbArr [3]Usb
	usbArr[0] = Phone{"小米"}
	usbArr[1] = Phone{"vivo"}
	usbArr[2] = Camera{"尼康"}

	fmt.Println(usbArr)

	//Phone 还有一个特有的方法 call()，请遍历 Usb 数组，如果是 Phone 变量，
	var computer Computer
	for _, val := range(usbArr) {
		computer.Working(val)
		fmt.Println()
	}

}

//细节说明