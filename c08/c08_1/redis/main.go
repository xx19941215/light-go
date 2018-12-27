package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	defer conn.Close()

	//Redis String的读写
	_, err = conn.Do("Set", "name", "xiaoxiao")
	if err != nil {
		fmt.Println("set  err=", err)
		return
	}

	v, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get  err=", err)
		return
	}

	fmt.Printf("操作ok %T\n", v)

	//Redis HASH的读写
	_, err = conn.Do("HSet", "nihao", "zh-cn", "你好")
	if err != nil {
		fmt.Println("HSet  err=", err)
		return
	}

	_, err = conn.Do("HSet", "nihao", "en-us", "hello")
	if err != nil {
		fmt.Println("HSet err=", err)
		return
	}

	r1, err := redis.String(conn.Do("HGet", "nihao", "zh-cn"))
	if err != nil {
		fmt.Println("HGet err=", err)
		return
	}

	r2, err := redis.String(conn.Do("HGet", "nihao", "en-us"))
	if err != nil {
		fmt.Println("HGet err=", err)
		return
	}

	fmt.Println("操作ok nihao zh-cn", r1)
	fmt.Println("操作ok nihao en-us", r2)

	//Redis HASH HMSET HMGET
	_, err = conn.Do("HMSet", "user", "name", "john", "age", 19)
	if err != nil {
		fmt.Println("HMSet  err=", err)
		return
	}

	r4, err := redis.Strings(conn.Do("HMGet", "user", "name", "john", "age"))
	if err != nil {
		fmt.Println("HMGet err=", err)
		return
	}

	fmt.Printf("操作ok user name john age = %T", r4)
}
