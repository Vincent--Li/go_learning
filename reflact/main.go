package main

import (
	"fmt"
	"reflect"
)

func main (){
	//1. 先定一个int
	var num int = 100
	reflectTest(num)
}

func reflectTest(a interface{}){
	// 通过反射获取的传入的变量 type, kind, 值
	// 1. 先获取到reflect.Type
	rType := reflect.TypeOf(a)
	fmt.Printf("%T, %v\n", rType, rType)

	// 2. 获取reflect.Value
	rVal := reflect.ValueOf(a)
	fmt.Printf("%T, %v\n", rVal	, rVal)

	// 3. kind 类别
	fmt.Printf("%s, %s\n", rType.Kind(), rVal.Kind())

	value := rVal.Int()
	fmt.Printf("%T, %v\n", value, value)

	// 3. 转成interface
	iV := rVal.Interface()
	fmt.Printf("%T, %v\n", iV, iV)

	// 4. interface 通过类型断言转回原类型
	value2 := iV.(int)
	fmt.Printf("%T, %v\n", value2, value2)
}