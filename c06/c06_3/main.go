package main

import (
	"fmt"
)

//channle本质就是一个数据结构-队列
//数据是先进先出【FIFO : first in first out】
//线程安全，多 goroutine 访问时，不需要加锁，就是说 channel 本身就是线程安全的
//channel有类型的，一个string的channel只能存放string类型数据。
type Person struct{}
var (
	intChan chan int
	mapChan chan map[int]string
	personChan chan Person
	personChan2 chan *Person
)

type Cat struct {
	Name string
	Age int
}

func main() {
	//创建一个可以存放 3 个 int 类型的管道
	intChan = make(chan int, 3)
	fmt.Printf("intChan 的值=%v intChan 本身的地址=%p\n", intChan, &intChan)

	//写数据
	intChan<- 20
	num := 211

	intChan<-num
	intChan<-50

	//看看管道的长度和 cap(容量)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan))

	//读数据
	var num2 int 
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan))

	//在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	num3 := <-intChan 
	num4 := <-intChan 
	//num5 := <-intChan
	fmt.Println("num3=", num3, "num4=", num4/*, "num5=", num5*/)

	//channel中只能存放指定的数据类型
	//channle的数据放满后，就不能再放入了
	//如果从 channel 取出数据后，可以继续放入
	//在没有使用协程的情况下，如果 channel 数据取完了，再取，就会报 dead lock

	allChan := make(chan interface{}, 3)
	allChan<- 3
	allChan<- "jack"
	cat := Cat{"xiao", 20,}
	allChan<-cat

	<-allChan
	<-allChan

	newCat := <-allChan
	fmt.Printf("newCat=%T newCat=%v", newCat, newCat)
	a := newCat.(Cat)//使用类型断言
	fmt.Printf("newCat.Name=%v", a.Name)

	//channel的遍历与关闭
	//使用内置函数 close 可以关闭 channel, 当 channel 关闭后，
	//就不能再向 channel 写数据了，但是仍然 可以从该 channel 读取数据	
	close(allChan)
	// allChan<-20

	//channel 的遍历
	//channel 支持 for--range 的方式进行遍历，注意两个细节
	//1.在遍历时，如果 channel 没有关闭，则回出现 deadlock 的错误
	//2.在遍历时，如果 channel 已经关闭，则会正常遍历数据，遍历完后，就会退出遍历。
	//3. for range 没有index 只有val
	
}

//细节说明
//1.channel 是引用类型
//2.channel 必须初始化才能写入数据, 即 make 后才能使用
//3.管道是有类型的，intChan 只能写入 整数 int