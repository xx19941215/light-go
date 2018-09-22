package main

import (
	"fmt"
)

//案例
func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

//writeData

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		//放入数据
		intChan <- i
		fmt.Println("writeData", i)
	}

	close(intChan)
}

//readData
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}

	exitChan <- true
	close(exitChan)
}

//细节说明
//如果只向管道写入数据，而没有读取，就会出现阻塞 而 deadlock。
