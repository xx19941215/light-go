package main

import (
	"time"
	"fmt"
	"sync"
)

var (
	myMap = make(map[int]int, 10)
	//因为没有对全局变量 m 加锁，因此会出现资源争夺问题，代码会出现错误，提示 concurrent map writes
	//解决方案:加入互斥锁
	lock sync.Mutex
)
//引入管道
func main() {
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//休眠10秒钟
	time.Sleep(time.Second * 10)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

//细节说明
//1.前面使用全局变量加锁同步来解决 goroutine 的通讯，但不完美
//2.主线程在等待所有 goroutine 全部完成的时间很难确定，我们这里设置 10 秒，仅仅是估算。
//3.如果主线程休眠时间长了，会加长等待时间，如果等待时间短了，可能还有 goroutine 处于工作状态，这时也会随主线程的退出而销毁
//4.通过全局变量加锁同步来实现通讯，也并不利用多个协程对全局变量的读写操作。
//5.上面种种分析都在呼唤一个新的通讯机制-channel