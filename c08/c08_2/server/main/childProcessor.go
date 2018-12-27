package main

import (
	"errors"
	"fmt"
	"light-go/c08/c08_2/common"
	"light-go/c08/c08_2/server/process"
	"light-go/c08/c08_2/server/utils"
	"net"
)

type ClientProcessor struct {
	conn net.Conn
	buf  [8192]byte
}

func (cp *ClientProcessor) ServerProcessMsg(msg *common.Message) (err error) {
	switch msg.Type {
	case common.LoginMesType:
		up := &process.UserProcess{
			Conn: cp.conn,
		}

		err = up.ServerProcessLogin(msg)

	case common.RegisterMesType:
		up := &process.UserProcess{
			Conn: cp.conn,
		}

		err = up.ServerProcessRegister(msg)

	default:
		err = errors.New("unsoupport message")
	}

	return
}

func (cp *ClientProcessor) Process() (err error) {
	for {
		var msg common.Message
		tf := &utils.Transfer{
			Conn: cp.conn,
		}

		msg, err = tf.ServerReadPackage()

		if err != nil {
			fmt.Println("server read package err=", err)
			return
		}

		err = cp.ServerProcessMsg(&msg)
		if err != nil {
			fmt.Println("process msg failed err=", err)
			return
		} else {
			return
		}

	}
}
