package main

import (
	"fmt"
	"time"
)

//细节说明
func main() {
	//1.channel可以声明为只读，或者只写性质

	//只可以写
	// var chan2 chan<- int

	// chan2 = make(chan int, 3)
	// chan2 <- 20
	// fmt.Println("chan2=", chan2)

	//声明只可以读
	// var chan3 <-chan int
	// num2 := <-chan3

	// fmt.Println("num2=", num2)
	//默认双向
	var ch chan int
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)

	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}

	fmt.Println("结束")

	//使用 select 可以解决从管道取数据的阻塞问题
	//1.定义一个管道 10 个数据
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//2.定义一个管道 5 个数据 string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock
	//可以使用 select 方式可以解决
	for {
		select {
		//注意: 这里，如果 intChan 一直没有关闭，不会一直阻塞而 deadlock //，会自动到下一个 case 匹配
		case v := <-intChan:
			fmt.Printf("从 intChan 读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从 stringChan 读取的数据%s\n", v)
			time.Sleep(time.Second)
		default:
			return
		}
	}

}

func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	var a struct{}

	exitChan <- a
}

func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch

		if !ok {
			break
		}

		fmt.Println(v)

	}
	var a struct{}
	exitChan <- a
}
