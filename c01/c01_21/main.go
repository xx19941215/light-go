package main

import (
	"time"
	"fmt"
	"math/rand"
)

//跳转控制
func main() {
	
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1 //生成[0, 100)
		fmt.Println("n=", n)
		count++
		if n == 99 {
			break
		}
	}

	fmt.Printf("生成99用了%d次", count)
	//break细节
	//break语句出现在多层嵌套的语句块中时，可以通过标签指明要终止的是哪一层语句块
	//break默认会跳出最近的for循环

	for1:for {
		for {
			break for1
		}
	}

	//continue
	//continue 语句用于结束本次循环，继续执行下一次循环。
	//continue 语句出现在多层嵌套的循环语句体中时，
	//可以通过标签指明要跳过的是哪一层循环, 这 个和前面的 break 标签的使用的规则一样.

	//goto
	//Go 语言的 goto 语句可以无条件地转移到程序中指定的行
	//goto语句通常与条件语句配合使用。可用来实现条件转移，跳出循环体等功能。
	//在 Go 程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序
	//都产生困难

	var n int = 30
	fmt.Println("ok1")
	if n > 20 {
		goto label1
	}

	fmt.Println("ok1")
	fmt.Println("ok1")
	fmt.Println("ok1")
	fmt.Println("ok1")
	
	label1:
	fmt.Println("ok2")
	fmt.Println("ok2")
	fmt.Println("ok2")

	//return
	//如果 return 是在普通的函数，则表示跳出该函数，即不再执行函数中 return 后面代码，也可以 理解成终止函数。
	//如果 return 是在 main 函数，表示终止 main 函数，也就是说终止程序。
}

//细节说明