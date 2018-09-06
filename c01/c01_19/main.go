package main

import (
	"fmt"
)

//流程控制
func main() {
	var age int = 40
	//顺序控制注意事项
	var age2 int = age + 10
	//必须采用合法的向前引用
	fmt.Println("age2=", age2)
	//分支控制
	if age2 > 20 {
		fmt.Println("age2>=20")
	}
	//golang在if判断中直接定义变量
	if age3 := 40; age3 > 20 {
		fmt.Println("age3>=20")
	}

	//双分支 if else
	//注意事项
	// {}不能省略  else 不能换行
	// if 后面的括号不需要
	//多分支控制
	// if else if else(不是必须)

	if age2 > 10 {
		fmt.Println("age2>=10")
	} else if age2 > 20 {
		fmt.Println("age2>=20")
	} else {
		fmt.Println("age=", age2)
	}

	//嵌套分支
	//建议最多三层
	//switch分支控制
	var key byte
	fmt.Println("请输入一个字符串,a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)
	switch key {
		case 'a':
			fmt.Println("你好a")
		case 'b':
			fmt.Println("你好b")
		default:
			fmt.Println("输入有误")
	}

	//说明
	//1. switch的执行的流程是，先执行表达式，得到值，然后和case的表达式进行比较，如果相等，
	//就匹配到，然后执行对应的 case 的语句块，然后退出 switch 控制。
	//2. 如果switch的表达式的值没有和任何的case的表达式匹配成功，则执行default的语句块。执行
	//后退出 switch 的控制.
	//3. golang的case后的表达式可以有多个，使用逗号间隔.
	//4. golang中的case语句块不需要写break, 因为默认会有,即在默认情况下，当程序执行完 case 语
	 //句块后，就直接退出该 switch 控制结构。
	 //5.case/switch 后是一个表达式( 即:常量值、变量、一个有返回值的函数等都可以)
	 //6.case 后的各个表达式的值的数据类型，必须和 switch 的表达式数据类型一致
	 //7.case后面的表达式如果是常量值(字面量)，则要求不能重复
	 //8.default 语句不是必须的.
	 //9.switch 后也可以不带表达式，类似 if --else 分支来使用。
	 //10.switch 后也可以直接声明/定义一个变量，分号结束，不推荐。
	 //11.switch 穿透-fallthrough ，如果在 case 语句块后增加 fallthrough ,则会继续执行下一个 case，也 叫 switch 穿透

	 var num int = 10
	 switch num {
	 	case 10:
			fmt.Println("ok1")
			fallthrough
		case 12:
			fmt.Println("ok2")
		case 32:
			fmt.Println("ok2")
		default:
			fmt.Println("没有匹配到")
	 }

	 //12.Type Switch:switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际指向的 变量类型

	 var x interface{}
	 var y = 10.0
	 x = y
	 switch i := x.(type) {
		case nil:
			fmt.Printf("x的类型是 %T", i)
		case int:
			fmt.Printf("x的类型是int")
		case float64:
			fmt.Printf("x的类型是float64")
		case func(int) float64:
			fmt.Printf("x的类型是func(int)")
		case bool, string:
			fmt.Printf("x的类型是bool或者string")
		default:
			fmt.Printf("未知类型")
	//switch 和 if 的比较
	//如果判断的具体数值不多，而且符合整数、浮点数、字符、字符串这几种类型。建议使用 swtich 语句，简洁高效。
	//其他情况:对区间判断和结果为 bool 类型的判断，使用 if，if 的使用范围更广。
	 }
}