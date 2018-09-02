package main
//布尔类型
// 1. 布尔类型也叫 bool 类型，bool 类型数据只允许取值 true 和 false
// 2. bool类型占1个字节。
// 3. bool 类型适于逻辑运算，一般用于程序流程控制
import (
	"unsafe"
	"fmt"
)

func main() {
	var b = false
	fmt.Println("b=", b)
	fmt.Println("b占用的空间=", unsafe.Sizeof(b))
}