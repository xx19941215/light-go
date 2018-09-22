package main

import (
	"fmt"
	"time"
)

//goroutine中使用recover，解决协程中出现panic，导致程序崩溃问题
func main() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

func test() {
	//这里我们可以使用 defer + recover
	defer func() {
		//捕获 test 抛出的 panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()

	//定义了一个 map
	var myMap map[int]string
	myMap[0] = "golang" //error
}
