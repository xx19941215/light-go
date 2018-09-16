package main

import (
	"fmt"
)

//切片
func main() {
	//切片的英文是 slice
	//切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制。
	//切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度 len(slice)都一样。
	//切片的长度是可以变化的，因此切片是一个可以动态变化数组。
	//切片定义的基本语法:
	//var 切片名 []类型
	
	var intArr [5]int = [...]int{1,2,3,44,4}
	slice := intArr[1:3]
	//slice就是切片名
	//intArr[1:3]表示slice引用到intArr这个数组
	//引用intArr数组的起始下标是1, 结束是3但是不包含
	fmt.Println("intArr", intArr)
	fmt.Println("slice的元素是", slice)
	fmt.Println("slice的个数是", len(slice))
	fmt.Println("slice的容量是", cap(slice))

	//第二种方式
	//通过 make 来创建切片
	//基本语法:var 切片名 []type = make([]type, len, [cap])
	//参数说明: type: 就是数据类型 len : 大小 cap :指定切片容量，可选，如果你分配了cap,则要求 cap>=len.
	var slice2 []float64 = make([]float64, 5, 10)
	slice2[1] = 10
	slice2[3] = 20
	//对于切片 务必使用make
	fmt.Println("slice2", slice2)
	//说明
	//1.通过 make 方式创建切片可以指定切片的大小和容量
	//2.如果没有给切片的各个元素赋值，那么就会使用默认值[int , float=> 0string =>””bool =>false]
	//3.通过 make 方式创建的切片对应的数组是由 make 底层维护，对外不可见，即只能通过 slice 去访问各个元素.

	//方式三
	//定义一个切片，直接就指定具体数组，使用原理类似 make 的方式
	var slice3 []string = []string{"tom", "jack", "tony"}
	fmt.Println("slice3", slice3)
	fmt.Println("slice3", len(slice3))
	fmt.Println("slice3", cap(slice3))

	//方式 1 和方式 2 的区别
	//方式一是直接引用数组，这个数组是事先存在的，程序是可见的。
	//方式二是通过make来创建切片，make也会创建一个数组，是由切片在底层维护的，程序员是看不见的。

	//切片的遍历
	//for和for range
	//切片使用细节
	//1. 切片初始化时 var slice = arr[startIndex:endIndex]
	//说明:从 arr 数组下标为 startIndex，取到 下标为 endIndex 的元素(不含 arr[endIndex])。
	//2.切片初始化时，仍然不能越界。范围在 [0-len(arr)] 之间，但是可以动态增长. 
	//var slice = arr[0:end] 可以简写 var slice = arr[:end]
	//var slice = arr[start:len(arr)] 可以简写: var slice = arr[start:]
	//var slice = arr[0:len(arr)] 可以简写: var slice = arr[:]
	//3. cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素。
	//4. 切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者 make 一个空间供切片来使用
	//5.切片可以继续切片

	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice4 := arr[1:4]
	for i := 0; i< len(slice3); i++ {
		fmt.Printf("slice4[%v]=%v\n", i, slice4[i])
	}

	slice5 := slice4[1:2]
	slice5[0] = 100

	//因为arr slice4 slice5指向的数据空间是同一个，所以改一个值其他都变了
	fmt.Println("slice4", slice4)
	fmt.Println("slice5", slice5)
	fmt.Println("arr", arr)

	//6. 用 append 内置函数，可以对切片进行动态追加
	var slice6 []int = []int{100,200,300}
	slice6 = append(slice6, 400, 500, 600)

	fmt.Println("slice6", slice6)
	slice6 = append(slice6, slice6...)

	fmt.Println("slice6", slice6)
	//切片 append 操作的底层原理分析
	//切片 append 操作的本质就是对数组扩容
	//go 底层会创建一下新的数组 newArr(按照扩容后大小)
	//将 slice 原来包含的元素拷贝到新的数组 newArr
	//slice 重新引用到 newArr
	//注意 newArr 是在底层来维护的，程序员不可见
	
	//切片的拷贝操作

	var slice7 []int = []int{1,2,3,4,5,8}
	var slice8 = make([]int, 10)
	copy(slice8, slice7)

	fmt.Println("slice7", slice7)
	fmt.Println("slice8", slice8)
	//copy(para1, para2) 参数的数据类型是切片
	//按照上面的代码来看, slice7 和 slice8 的数据空间是独立，相互不影响，也就是说 slice7[0]= 999,slice8[0] 仍然是 1

	var a []int = []int{1,2,3,4}
	var slice9 = make([]int, 1)
	fmt.Println(slice9)
	copy(slice9, a)
	fmt.Println(slice9)
	
	//切片是引用类型，所以在传递时，遵守引用传递机制。
}


//细节说明
//1.slice 的确是一个引用类型
//2. slice 从底层来说，其实就是一个数据结构(struct 结构体)
//type slice struct {
//ptr *[2]int 
//len int 
//cap int
//}

