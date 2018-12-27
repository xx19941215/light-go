package process

import (
	"fmt"
)

type ClinetMgr struct {
	onlineUser map[int]*UserProcess
}

var (
	clinetMgr *ClinetMgr
)

func init() {
	clinetMgr = &ClinetMgr{
		onlineUser: make(map[int]*UserProcess, 1024),
	}
}

func (p *ClinetMgr) AddClient(userId int, client *UserProcess) {
	p.onlineUser[userId] = client
}

func (p *ClinetMgr) GetClient(userId int) (clinet *UserProcess, err error) {
	clinet, ok := p.onlineUser[userId]
	if !ok {
		fmt.Errorf("id=%d的用户不存在..", userId)
		return
	}

	return
}

func (p *ClinetMgr) GetAllUsers() map[int]*UserProcess {
	return p.onlineUser
}

func (p *ClinetMgr) DelClient(userId int) {
	delete(p.onlineUser, userId)
}
