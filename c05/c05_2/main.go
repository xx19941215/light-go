package main

import (
	"io"
	"fmt"
	"os"
	"bufio"
)

//文件
func main() {
	file, err := os.Open("/var/log/file.log")
	//打开文件
	//概念说明: file 的叫法
	//1. file 叫 file对象
	//2. file 叫 file指针
	//3. file 叫 file 文件句柄
	if err != nil {
		fmt.Println("open file error", err)
	}

	fmt.Printf("file=%v\n", file)

	//及时关闭file句柄，否则会有内存泄漏
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		// 读到一个换行就结束
		str, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		fmt.Printf(str)
	}

	fmt.Println("文件读取结束...")
	
}

//细节说明