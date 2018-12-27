package process

import (
	"encoding/json"
	"fmt"
	"light-go/c08/c08_2/common"
	"light-go/c08/c08_2/server/model"
	"light-go/c08/c08_2/server/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	userId int
	Buf    [8192]byte
}

func (up *UserProcess) NotifyOthersUserOnline(userId int) {
	users := clinetMgr.GetAllUsers()
	for id, client := range users {
		if id == userId {
			continue
		}

		client.NotifyUserOnline(id)
	}
}

func (up *UserProcess) NotifyUserOnline(userId int) {
	var respMsg common.Message
	respMsg.Type = common.UserStatusNotifyMesType
	var userStatusNotifyMes common.UserStatusNotifyMes

	userStatusNotifyMes.UserId = userId
	userStatusNotifyMes.Status = common.UserStatusOnline

	data, err := json.Marshal(userStatusNotifyMes)

	if err != nil {
		fmt.Println("Marshal failed", err)
		return
	}

	respMsg.Data = string(data)
	data, err = json.Marshal(respMsg)

	if err != nil {
		fmt.Println("Marshal failed", err)
		return
	}

	tf := &utils.Transfer{
		Conn: up.Conn,
	}

	err = tf.ServerWritePackage(data)

	if err != nil {
		fmt.Println("send msg failed", err)
		return
	}

	return
}

func (up *UserProcess) ServerProcessLogin(msg *common.Message) (err error) {
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(msg.Data), &loginMes)

	if err != nil {
		fmt.Println("unmarshal failed, err=", err)
	}

	var resMsg common.Message
	resMsg.Type = common.LoginMesType
	var loginResMes common.LoginResMes

	var user *common.User
	user, err = model.MyUserDao.Login(loginMes.Id, loginMes.Pwd)
	if err != nil {
		fmt.Println("登陆有错误, err=", err)
		loginResMes.Code = 500
		return
	} else {
		fmt.Println("登陆成功, 用户信息是", user)
		loginResMes.Code = 200
	}

	clinetMgr.AddClient(loginMes.Id, up)
	up.userId = loginMes.Id

	up.NotifyOthersUserOnline(loginMes.Id)

	userMap := clinetMgr.GetAllUsers()

	for userId, _ := range userMap {
		loginResMes.User = append(loginResMes.User, userId)
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("marshal err=", err)
		return
	}

	fmt.Println("返回的package包=%v", string(data))

	tf := &utils.Transfer{
		Conn: up.Conn,
	}

	err = tf.ServerWritePackage(data)

	if err != nil {
		fmt.Println("send msg err=", err)
		return
	}

	return

}

func (up *UserProcess) ServerProcessRegister(msg *common.Message) (err error) {
	var registerMes common.RegisterMes

	err = json.Unmarshal([]byte(msg.Data), &registerMes)

	if err != nil {
		fmt.Println("unmarshal failed err=", err)
	}

	var resMsg common.Message
	resMsg.Type = common.RegisterMesResType
	var registerResMes common.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		fmt.Printf("用户id已经被占用err=%v\n", err)
		registerResMes.Code = 100
		return
	} else {
		fmt.Println("没有错误 注册成功")
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("marshal err=", err)
		return
	}

	resMsg.Data = string(data)

	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: up.Conn,
	}

	err = tf.ServerWritePackage(data)

	if err != nil {
		fmt.Println("send msg err=", err)
		return
	}

	return

}
