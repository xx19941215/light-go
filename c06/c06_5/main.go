package main

import (
	"time"
	"fmt"
)

//go求素数

func putNum(intChan chan int) {
	for i := 1; i <= 80000; i ++ {
		intChan<- i
	}

	close(intChan)
}

//从intChan中取出数据，判断是否为素数，是的话就放到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		//取不到就是没有数了
		if !ok {
			break
		}

		flag = true
		//判断是否是素数
		for i := 2; i < num; i++ {
			if num % 2 == 0 {
				flag = false
				break
			}
		}

		if flag {
			primeChan<- num
		}
	}

	fmt.Println("一个协程退出")
	//这里还不能关闭primeChan
	exitChan<-true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool, 8)

	start := time.Now().Unix()

	go putNum(intChan)

	for i := 0; i < 8; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func(){
		for i := 0; i < 8; i++ {
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end - start)
		//8个协程都退出 关闭chan
		close(primeChan)
	}()

	for {
		_, ok := <-primeChan

		if !ok {
			break
		}

		// fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("主线程退出")

	test()

}

//细节说明

func test() {
	start := time.Now().Unix()
		for num := 1; num <= 80000; num++ {

			flag := true //假设是素数
			//判断num是不是素数
			for i := 2; i < num; i++ {
				if num % i == 0 {//说明该num不是素数
					flag = false
					break
				}
			}

			if flag {
				//将这个数就放入到primeChan
				//primeChan<- num
			}

		}
		end := time.Now().Unix()
		fmt.Println("普通的方法耗时=", end - start)
}