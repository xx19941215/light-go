package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"light-go/c08/c08_2/common"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8192]byte
}

func (transfer *Transfer) ServerReadPackage() (msg common.Message, err error) {
	n, err := transfer.Conn.Read(transfer.Buf[0:4])

	if n != 4 {
		err = errors.New("read head failed")
		return
	}

	fmt.Println("read package:", transfer.Buf[0:4])

	var packLen uint32
	packLen = binary.BigEndian.Uint32(transfer.Buf[0:4])

	fmt.Printf("receive len:%d\n", packLen)

	n, err = transfer.Conn.Read(transfer.Buf[0:packLen])

	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}

	fmt.Printf("receive data:%s\n", string(transfer.Buf[0:packLen]))
	err = json.Unmarshal(transfer.Buf[0:packLen], &msg)

	if err != nil {
		fmt.Println("unmarshal err=", err)
	}

	return
}

func (transfer *Transfer) ServerWritePackage(data []byte) (err error) {
	packLen := uint32(len(data))

	binary.BigEndian.PutUint32(transfer.Buf[0:4], packLen)

	n, err := transfer.Conn.Write(transfer.Buf[0:4])
	if err != nil {
		fmt.Println("write data failed")
		return
	}

	n, err = transfer.Conn.Write(data)
	if err != nil {
		fmt.Println("write data failed")
		return
	}

	if n != int(packLen) {
		fmt.Println("write data failed")
		err = errors.New("write data failed")
		return
	}

	return
}
