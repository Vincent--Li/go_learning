package main

import (
	"fmt"
	"strconv"
)

func main(){

	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("%T %v\n", b, b)

	var str2 string = "1234590"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 0, 0)
	fmt.Printf("%T %v\n", n1, n1)
}
