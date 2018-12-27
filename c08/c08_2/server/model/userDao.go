package model

import (
	"encoding/json"
	"fmt"
	"light-go/c08/c08_2/common"

	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserMgr(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}

	return
}

func (ud *UserDao) getUser(conn redis.Conn, id int) (user *common.User, err error) {
	result, err := redis.String(conn.Do("HGet", "users", fmt.Sprintf("%d", id)))

	fmt.Printf("result=%v\n", result)

	if err != nil {
		if err == redis.ErrNil {
			err = ErrUserNotExist
		}

		return
	}

	user = &common.User{}

	err = json.Unmarshal([]byte(result), user)

	if err != nil {
		fmt.Println("err=", err)
		return
	}

	fmt.Printf("从数据库得到的user=%v", user)

	return

}

func (ud *UserDao) Register(user *common.User) (err error) {
	conn := ud.pool.Get()
	defer conn.Close()

	_, err = ud.getUser(conn, user.Id)
	if err == nil {
		err = ErrUserExist
		return
	}

	data, err := json.Marshal(user)

	if err != nil {
		return
	}

	_, err = conn.Do("Hset", "users", fmt.Sprintf("%d", user.Id), string(data))

	if err != nil {
		return
	}

	return

}

func (ud *UserDao) Login(id int, passwd string) (user *common.User, err error) {
	conn := ud.pool.Get()
	defer conn.Close()

	user, err = ud.getUser(conn, id)

	if err != nil {
		return
	}

	if user.Pwd != passwd {
		err = ErrInvalidPasswd
		return
	}

	return
}
