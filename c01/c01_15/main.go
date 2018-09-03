package main

import (
	"fmt"
)

//逻辑运算符
func main() {
	var age int = 40
	if (age < 50 || age > 30) {
		fmt.Println("so young")
	}
}

//细节说明
//&&也叫短路与:如果第一个条件为false，则第二个条件不会判断，最终结果为false
//||也叫短路或:如果第一个条件为true，则第二个条件不会判断，最终结果为true