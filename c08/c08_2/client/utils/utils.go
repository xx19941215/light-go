package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"light-go/c08/c08_2/common"
	"net"
	"time"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8192]byte
}

func (tf *Transfer) ClientReadPackage() (msg common.Message, err error) {
	n, err := tf.Conn.Read(tf.Buf[0:4])

	if n != 4 {
		err = errors.New("read header failed")
		fmt.Println("如果读取错误，则休息三十秒，再发送数据...")
		time.Sleep(time.Second * 30)
		return
	}

	var packLen uint32
	packLen = binary.BigEndian.Uint32(tf.Buf[0:4])

	n, err = tf.Conn.Read(tf.Buf[0:packLen])

	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}

	err = json.Unmarshal(tf.Buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed err=", err)
	}

	return
}
