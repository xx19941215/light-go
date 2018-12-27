package process

import (
	"encoding/json"
	"fmt"
	"light-go/c08/c08_2/client/utils"
	"light-go/c08/c08_2/common"
	"net"
)

func ProcessServerMessage(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	for {
		msg, err := tf.ClientReadPackage()

		if err != nil {
			fmt.Println("read err:", err)
			fmt.Println("msg=", msg, "err=", err)
		}

		var userStatus common.UserStatusNotifyMes
		err = json.Unmarshal([]byte(msg.Data), &userStatus)

		if err != nil {
			fmt.Println("unmarshal failed, err:", err)
			return
		}

		switch msg.Type {
		case common.UserStatusNotifyMesType:

		}
	}
}
