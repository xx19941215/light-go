package main

import (
	"fmt"
	"light-go/c08/c08_2/server/model"
	"net"
	"time"
)

func initUserDao() {
	model.MyUserDao = model.NewUserMgr(pool)
}

func processMain(conn net.Conn) {
	cp := &ClientProcessor{
		conn: conn,
	}

	err := cp.Process()

	if err != nil {
		fmt.Println("client process failed err=", err)
		return
	}
}

func main() {
	initRedis("localhost:6379", 16, 1024, time.Second*300)

	initUserDao()

	var addr = "0.0.0.0:8889"
	fmt.Println("服务器在8889端口监听，等待客户端链接")

	l, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Println("监听失败, err=", err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept failed, err=", err)
			continue
		}

		go processMain(conn)
	}
}
