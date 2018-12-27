package common

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterMesResType      = "RegisterMesRes"
	UserStatusNotifyMesType = "UserStatusNotifyMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	Id   int    `json:"userId"`
	Name string `json:"userName"`
	Pwd  string `json:"userPwd"`
}

type LoginResMes struct {
	Code  int    `json:"Code"`
	User  []int  `json:"users"`
	Error string `json:"error"`
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type UserStatusNotifyMes struct {
	UserId int `json:"userId"`
	Status int `json:"userStatus"`
}
