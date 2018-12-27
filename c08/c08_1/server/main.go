package main

import (
	"fmt"
	_ "io"
	"net"
)

func main() {
	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")

	if err != nil {
		fmt.Println("listen err=", err)
		return
	}

	defer listen.Close()

	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept suc con=%v 客户端ip为=%v\n", conn, conn.RemoteAddr().String())
		}

		go process(conn)

	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("客户端退出err=%v", err)
			return
		}

		fmt.Print(string(buf[:n]))
	}
}
