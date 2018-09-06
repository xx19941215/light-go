package main

import (
	"fmt"
)

//二维数组
func main() {
	//使用方式1
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2 
	arr[2][3] = 3 

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}

		fmt.Println()
	}

	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println(arr2)

	fmt.Printf("arr2[0]的地址%p\n", &arr2[0])
	fmt.Printf("arr2[1]的地址%p\n", &arr2[1])


	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][0])
	fmt.Printf("arr2[1][0]的地址%p\n", &arr2[1][0])

	//使用方式二
	//直接初始化
	//var 数组名 [大小][大小]类型 = [大小][大小]类型{{初值..},{初值..}}
	var arr3 [2][3]int = [2][3]int{{1,2,3}, {4, 5, 6}}
	fmt.Print(arr3)
	//二维数组在声明/定义时也对应有四种写法[和一维数组类似]
	//var 数组名 [大小][大小]类型 = [大小][大小]类型{{初值..},{初值..}}
	//var 数组名 [大小][大小]类型 = [...][大小]类型{{初值..},{初值..}}
	//var 数组名 = [大小][大小]类型{{初值..},{初值..}}
	//var 数组名 = [...][大小]类型{{初值..},{初值..}}

	//二维数组的遍历
	//双层 for 循环完成遍历
	//for-range方式完成遍历
}


//细节说明