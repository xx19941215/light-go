package main
//数据类型
//1. 基本数据类型
// 1.1 数值型 包括整数类型(int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 byte)、浮点类型(float32 float64)
// 1.2 字符型 没有专门的字符型 使用byte来来保存单个字母字符
// 1.3 布尔型(bool)
// 1.4 字符串(string)
//2. 派生/复杂数据类型 包括指针、数组、结构体、管道、函数、切片、接口、Map
import (
	"unsafe"
	"fmt"
)
func main() {
	var i int = 1
	fmt.Println("i=", i)	

	var j int8 = 127
	fmt.Println("i=", j)	

	// int 有符号 32位系统占用4字节 可以表示 负2的31次方到 2的31次方-1 64位占用 8字节 可以表示负2的63次方到2的63次方-1
	// uint 无符号 32位系统占用4字节 可以表示 0到2的32次方-1 64位占用 8字节 可以表示0到2的64次方-1
	// rune 有符号 位用 4字节 可以表示 负2的31次方到 2的31次方-1 等价int32表示一个unicode编码
	// byte 无符号 与unit8等价 表示0-255 当要存储字符的时候用byte

	var a int = 233
	fmt.Println("a=", a)
	var b uint = 233233
	fmt.Println("b=", b)
	var c byte = 255
	fmt.Println("c=", c)

	// 使用细节
	// 1. Go各整数类型分有无符号。int uint 的大小和系统有关
	// 2. Go的整型默认声明为int型

	var n1 = 100
	fmt.Printf("n1的类型%T \n", n1)
	// 查看某个变量占用的字节大小和数据类型
	var n2 int64 = 10
	fmt.Printf("n2的类型是 %T n2占用的字节数是%d", n2, unsafe.Sizeof(n2))
	// Go程序中整型变量在使用时，遵循保小不保大的原则。即再保证程序运行正常的情况下，尽量使用占用空间小的数据类型
}