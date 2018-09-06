package main

import (
	"strings"
	"strconv"
	"fmt"
)

//字符串函数
func main() {
	//统计字符串的长度，按字节 len(str)
	str := "hello北"
	fmt.Println("str len=", len(str))
	//避免乱码
	str2 := "hello肖"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符:%c\n", r[i])
	}
	//字符串转整数
	n, err := strconv.Atoi("12")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Printf("数字:%d\n", n)
	}
	//整数转字符串
	str3 := strconv.Itoa(123445)
	fmt.Printf("str=%v, str=%T\n", str3, str3)

	//字符串转[]byte 
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	//byte转字符串
	str3 = string([]byte{97, 98, 99})	
	fmt.Printf("str=%v, str=%T\n", str3, str3)

	//10进制转2，8，6进制
	num := strconv.FormatInt(123,2)
	fmt.Printf("num=%v, num=%T\n", num, num)

	//查找子串是否在指定的字符串中: strings.Contains("seafood", "foo") //true
	var isHas bool
	isHas = strings.Contains("seafood", "foo")

	fmt.Printf("isHas=%v, isHas=%T\n", isHas, isHas)

	//统计一个字符串有几个指定的子串 : strings.Count("ceheese", "e") //4

	//不区分大小写的字符串比较(==是区分字母大小写的): fmt.Println(strings.EqualFold("abc", "Abc")) // true

	//返回子串在字符串第一次出现的 index 值，如果没有返回-1 : strings.Index("NLT_abc", "abc") // 4

	//返回子串在字符串最后一次出现的 index，如没有返回-1 : strings.LastIndex("go golang", "go")

	//将指定的子串替换成 另外一个子串: strings.Replace("go go hello", "go", "go 语言", n) 
	//n 可以指 定你希望替换几个，如果 n=-1 表示全部替换

	//按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组:
	strings.Split("hello,wrold,ok", ",")

	//将字符串的字母进行大小写的转换: strings.ToLower("Go") // go strings.ToUpper("Go") // GO

	//将字符串左右两边的空格去掉: strings.TrimSpace(" tn a lone gopher ntrn ")

	//将字符串左右两边指定的字符去掉 : strings.Trim("! hello! ", " !") 和 " "去掉

	//将字符串左边指定的字符去掉 : strings.TrimLeft("! hello! ", " !") "去掉

	//将字符串右边指定的字符去掉 : strings.TrimRight("! hello! ", " !") "去掉

	//判断字符串是否以指定的字符串开头: strings.HasPrefix("ftp://192.168.10.1", "ftp") // true

	//判断字符串是否以指定的字符串结束: strings.HasSuffix("NLT_abc.jpg", "abc") //false

}

//细节说明
//&&也叫短路与:如果第一个条件为false，则第二个条件不会判断，最终结果为false
//||也叫短路或:如果第一个条件为true，则第二个条件不会判断，最终结果为true