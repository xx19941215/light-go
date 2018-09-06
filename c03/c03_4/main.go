package main

import (
	"fmt"
)

//string和slice
func main() {
	//string底层是一个byte数组，因此string也可以进行切片处理
	str := "hello world"
	slice := str[6:]
	fmt.Println("slice", slice)
	//string 和切片在内存的形式，内存示意图见图片
	//string 是不可变的，也就说不能通过 str[0] = 'z' 方式来修改字符串
	//如果需要修改字符串，可以先将 string -> []byte / 或者 []rune -> 修改 -> 重写转成 string
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println(str)

	//细节 我们转成byte之后，可以处理英文和数字，但是不能处理汉字
	//原因是byte是按照字节处理的，汉字会出现乱码
	//解决方法是将string转成[]rune即可，因为[]rune是按照字符处理的，兼容汉字
	arr3 := []rune(str)
	arr3[0] = '北'
	str = string(arr3)
	fmt.Println(str)
}

//可以接收一个 n int
//能够将斐波那契的数列放到切片中

func fbn(n int) ([]uint64){
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[n - 1] + fbnSlice[n - 2]
	}

	return fbnSlice
}