package main

import (
	"fmt"
)

//map
func main() {
	//map 是 key-value 数据结构，又称为字段或者关联数组。类似其它编程语言的集合，
	//var map 变量名 map[keytype]valuetype
	//key 可以是什么类型?
	//golang 中的 map，的 key 可以是很多种类型，比如 bool, 数字，string, 指针, channel , 
	//还可以是只 包含前面几个类型的 接口, 结构体, 数组。通常 key 为 int 、string
	//注意: slice， map 还有 function 不可以，因为这几个没法用 == 来判断
	//valuetype 可以是什么类型？
	//valuetype 的类型和 key 基本一样，这里就不再赘述了
	//通常为: 数字(整数,浮点数),string,map,struct
	//map声明的举例:
	//var a map[string]string
	//var a map[string]string
	//var a map[int]string
	//var a map[string]map[string]string
	//注意:声明是不会分配内存的，初始化需要 make ，分配内存后才能赋值和使用。

	var a map[string]string
	//使用map前，需要make,make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "hello"
	a["no2"] = "world"
	fmt.Println(a)

	//细节
	//1) map在使用前一定要make
	//2) map的key是不能重复，如果重复了，则以最后这个key-value为准
	//3) map的value是可以相同的.
	//4) map 的 key-value 是无序
	//5) make内置函数数目

	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"

	fmt.Println(cities)
	//第三种方式

	heros := map[string]string{
		"hero1" : "宋江",
		"hero2" : "卢俊义",
	}

	fmt.Println(heros)

	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tome"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["age"] = "17"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "jenny"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["age"] = "18"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"])
	fmt.Println(studentMap["stu02"]["age"])

	//map crud
	//map增加和更新:
	//map["key"] = value //如果 key 还没有，就是增加，如果 key 存在就是修改。
	//map删除:
	//delete(map，"key") ，delete 是一个内置函数，如果 key 存在，
	//就删除该 key-value,如果 key 不存在， 不操作，但是也不会报错
	delete(studentMap, "stu01")
	fmt.Println(studentMap)

	//细节说明
	//如果我们要删除 map 的所有 key ,没有一个专门的方法一次删除，可以遍历一下 key, 
	//逐个删除 或者 map = make(...)，make 一个新的，让原来的成为垃圾，被 gc 回收

	// cities = make(map[string]string)
	// fmt.Println(cities)

	//map查找:
	val, ok := cities["no2"]
	if ok {
		fmt.Printf("有no1 key值为%v\n", val)
	} else {
		fmt.Printf("没有no1 key值\n")
	}

	//如果 cities 这个 map 中存在 "no2" ， 那么 ok 就会返回 true,否则返回 false

	//map遍历 for range 不解释

}

//细节说明