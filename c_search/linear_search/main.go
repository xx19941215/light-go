package main

import (
	"fmt"
)

//
func main() {
	names := [4]string{"tom", "tony", "xiao", "jenny"}
	var search string
	fmt.Println("请输入要查找的name")
	fmt.Scanln(&search)

	// for i := 0; i < len(names); i++ {
	// 	if (names[i] == search) {
	// 		fmt.Printf("找到了%v, 下标为%v", names[i], i)
	// 		break
	// 	} else if (i == len(names) - 1) {
	// 		fmt.Printf("没有找到%v", search)
	// 	}
	// }

	linearSearch(&names, search)
}

//第二种张方式

func linearSearch(arr *[4]string, search string) {
	index := -1
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == search {
			index = i
			break
		}
	}

	if index != -1 {
		fmt.Printf("找到了%v, 下标为%v", (*arr)[index], index)
	} else {
		fmt.Printf("没有找到%v", search)
	}
}