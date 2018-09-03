package main

//值类型和引用类型
//值类型包括: 基本数据类型 int 系列, float 系列, bool, string 、数组和结构体 struct
//引用类型: 指针、slice 切片、map、管道 chan、interface 等都是引用类型
func main() {
	//值类型:变量直接存储值，内存通常在栈中分配
	//引用类型:变量存储的是一个地址，这个地址对应的空间才真正存储数据(值)，
	//内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由 GC 来回收
}

//标识符
//1.Golang 对各种变量、方法、函数等命名时使用的字符序列称为标识符
//2.凡是自己可以起名字的地方都叫标识符
//命名规则
//1.由 26 个英文字母大小写，0-9 ，_ 组成
//2.数字不可以开头。var num int //ok var 3num int //error
//3.Golang中严格区分大小写。
//4.标识符不能包含空格
//5.下划线"_"本身在 Go 中是一个特殊的标识符，称为空标识符。
//可以代表任何其它的标识符，但 是它对应的值会被忽略(比如:忽略某个返回值)。所以仅能被作为占位符使用，不能作为标识符使用
//6.不能以系统保留关键字作为标识符(一共有 25 个)，比如 break，if 等等...

//标识符注意事项
//1.包名:保持 package 的名字和目录保持一致，尽量采取有意义的包名，简短，有意义，不要和 标准库不要冲突 fmt
//2.变量名、函数名、常量名:采用驼峰法
//3.如果变量名、函数名、常量名首字母大写，则可以被其他的包访问;如果首字母小写，则只能
//在本包中使用 ( 注:可以简单的理解成，首字母大写是公开的，首字母小写是私有的) 
//在 golang 没有 public , private 等关键字。

//系统保留关键字
//break default func interface select case defer go map struct chan else goto package switch
//const fallthrough if range type continue for import return var

//系统预定义标识符
//append bool byte cap close complex complex64 complex128 uint16 copy false float32 float64 imag
//int8 int16 uint32 int32 int64 iota len make new nil panic uint64 print println real recover
//string true uint unit8 uintprt 