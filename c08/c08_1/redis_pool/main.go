package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Rpush", "list", "xiao")

	if err != nil {
		fmt.Println("Rpush err=", err)
		return
	}

	r, err := redis.String(conn.Do("Rpop", "list"))

	if err != nil {
		fmt.Println("Rpop err=", err)
		return
	}

	fmt.Printf("list rpop data=%v\n", r)

	conn2 := pool.Get()

	_, err = conn2.Do("Zadd", "zaddList", 1, "one")

	if err != nil {
		fmt.Println("Zadd err=", err)
		return
	}

	r1, err := redis.Int64(conn.Do("Zcard", "zaddList"))

	if err != nil {
		fmt.Println("Zcard err=", err)
		return
	}

	fmt.Printf("list length data=%v\n", r1)
}
