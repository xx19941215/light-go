// Go变量使用的三种方式
package main
import "fmt"

//定义全局变量
var n1 = 100
var n2 = 200
var name = "xiao"

//也可以一次性声明

var (
	n3 = 100
	n4 = 200
	name2 = "xiao"
)

func main() {
	//1. 指定变量类型 声明后若不赋值，使用默认值int的默认值是0
	var i int
	fmt.Println("i=", i)	
	//2. 根据值自动判断类型 (类型推导)
	var num = 10.11
	fmt.Println("num=", num)
	//3. 省略var 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误
	name := "tom"
	fmt.Println("name=", name)
	//4. 一次性声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
	//5. 第二种一次性声明多个变量
	var n4, name2, n5 = 100, "tom", 888
	fmt.Println("n4=", n4, "name2=", name2, "n5=", n5)
	//6. 一次性声明多个变量同样可以使用类型推导
	n6, name3 := 200, "xiaoxiao"
	fmt.Println("n6=", n6, "name3=", name3)
	//7. 该区域内的数据值可以在范围内不断变化(不能改变数据类型)
	n4 = 200
	n4 = 300
	fmt.Println("n4=", n4)
	//8. 变量在一个函数或者代码块中不能重名
	//9. 变量 = 变量名 + 值 + 数据类型
	//10.变量如果没有赋予初始值，int默认值为0 string默认为空串 小数默认为0
	//11.关于+ 当变量都是数值型的时候做加法操作，当都是字符串的时候作字符串连接操作
}