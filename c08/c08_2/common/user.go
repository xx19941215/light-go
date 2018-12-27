package common

const (
	UserStatusOnline  = 1
	UserStatusOffline = iota
)

type User struct {
	Id     int    `json:"userId"`
	Name   string `json:"userName"`
	Pwd    string `json:"userPwd"`
	Status int    `json:"status"`
}
