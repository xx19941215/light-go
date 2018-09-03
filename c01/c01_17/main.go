package main

import (
	"fmt"
)

//键盘输入语句
func main() {
	//要求:可以从控制台接收用户信息，【姓名，年龄，薪水, 是否通过考试 】。
	var (
		name string
		age byte
		salary float64
		isPass bool
	)

	// fmt.Println("请输入姓名")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水")
	// fmt.Scanln(&salary)
	// fmt.Println("是否通过考试")
	// fmt.Scanln(&isPass)

	//第二种方式
	fmt.Println("请输入姓名 年龄 薪水 是否通过考试 按空格隔开");
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)
	fmt.Printf("姓名：%v\t年龄：%v\t薪水：%v\t是否通过考试：%v\t", name, age, salary, isPass);
}
