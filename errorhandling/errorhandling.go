package main

import (
	"fmt"
	"time"
)

func test()	{

	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println("res=", result)

}

func main(){
	test()
	fmt.Println("这里执行不到")

	time.Sleep(time.Second)
}
