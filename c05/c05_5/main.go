package main

import (
	"flag"
	"fmt"
)

//flag包解析命令行参数
func main() {
	//基础方式
	// fmt.Println("命令行参数有", len(os.Args))

	// for i,v := range(os.Args) {
	// 	fmt.Printf("args[%v]=%v\n", i, v)
	// }

	var user string
	var pwd string
	var host string
	var port int

	//&user就是接受用户命令行中输入的 -u 后面的参数值
	//“”默认值
	//u 就是-u指定参数
	//"用户名，默认为空" 说明

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "用户密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口，默认为3306")

	//必须调用下
	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}

//细节说明