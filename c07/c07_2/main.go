package main

import (
	"fmt"
	"reflect"
)

//
func main() {
	question01()
	question02()
	question03()
}

func question02() {
	var str string = "tom"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jack")
	fmt.Printf("%v", str)
}

func question01() {
	var v float64 = 1.2
	rVal := reflect.ValueOf(v)

	rType := reflect.TypeOf(v)

	iV := rVal.Interface()
	v2 := iV.(float64)

	kind1 := rVal.Kind()
	kind2 := rType.Kind()

	fmt.Println("v2=", v2)
	fmt.Printf("kind1=%v kind2=%v", kind1, kind2)
}

type Cal struct {
	Num1 int
	Num2 int
}

func (this Cal) GetSub(name string) string {
	res := this.Num1 - this.Num2
	return fmt.Sprintf("%v完成了减法运行,%v - %v = %v", name, this.Num1, this.Num2, res)
}

func question03() {
	a := Cal{
		Num1: 8,
		Num2: 5,
	}

	val := reflect.ValueOf(a)

	num := val.NumField()

	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
	}

	var params []reflect.Value
	params = append(params, reflect.ValueOf("tom"))

	res := val.Method(0).Call(params)

	fmt.Println(res[0].String())
}

//细节说明
