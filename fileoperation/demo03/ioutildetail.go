package main

import (
	"fmt"
	"io/ioutil"
)

func main(){

	// 使用ioutil一次性读取到位
	file := "../test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(content)

	fmt.Printf("%s", content)

	fmt.Println(string(content))
}
