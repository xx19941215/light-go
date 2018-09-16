package main

import (
	"encoding/json"
	"fmt"
)

//json序列化
//json 序列化是指，将有 key-value 结构的数据类型(比如结构体、map、切片)序列化成 json 字符串的操作。
//这里我们介绍一下结构体、map 和切片的序列化，其它数据类型的序列化类似。

//对于结构体的序列化，如果我们希望序列化后的 key 的名字，又我们自己重新制定，那么可以给 struct
//指定一个 tag 标签.
type Monster struct {
	Name string `json:"monster_name"`
	Age int
	Birthday string
	Sal string
	Skill string
}

func main() {
	monster := Monster{
		Name: "悟空",
		Birthday: "1994-12-15",
		Age: 20,
		Sal: "1000",
		Skill: "coding",
	}

	data, err := json.Marshal(&monster)
	//反序列化
	//json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v", err)
	}

	fmt.Printf("序列化后=%v\n", string(data))
	testMap()

	testSlice()
}

func testMap() {
	//key是string value 为任意类型
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "xiaoxiao"
	a["age"] = 20

	//map本身就是引用类型
	data, err := json.Marshal(a)
	//反序列化
	//注意:反序列化 map,不需要 make,因为 make 操作被封装到 Unmarshal 函数 
	//err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v", err)
	}

	fmt.Printf("序列化后=%v\n", string(data))
}

func testSlice() {
	var slice []map[string]interface{}	
	var m1 map[string]interface{}

	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 20
	slice = append(slice, m1)

	var m2 map[string]interface{}

	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 21
	slice = append(slice, m2)

	data, err := json.Marshal(slice)

	//反序列化，不需要 make,因为 make 操作被封装到 Unmarshal 函数 
	//err := json.Unmarshal([]byte(str), &slice)

	if err != nil {
		fmt.Printf("序列化错误 err=%v", err)
	}

	fmt.Printf("序列化后=%v\n", string(data))
}

//在反序列化一个json字符串时，要确保反序列化后的数据类型和原来序列化前的数据类型一致。
//如果 json 字符串是通过程序获取到的，则不需要再对 “ 转义处理。