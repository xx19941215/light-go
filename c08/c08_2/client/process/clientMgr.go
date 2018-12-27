package process

var onlineUserMap map[int]*common.User = make(map[int]*common.User, 10)
var UserId int

func output