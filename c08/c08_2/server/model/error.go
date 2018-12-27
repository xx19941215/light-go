package model

import "errors"

var (
	ErrUserNotExist  = errors.New("user not exist")
	ErrInvalidPasswd = errors.New("username or passwd not correct")
	ErrInvalidParams = errors.New("Invalid params")
	ErrUserExist     = errors.New("user exist")
)
