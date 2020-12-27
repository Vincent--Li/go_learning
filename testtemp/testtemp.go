package main

import (
	"fmt"
	"go.learning/m/v2/tempconv"
)

var a, b, c int

func main(){
	fmt.Printf(tempconv.AbsoluteZeroC.String())
	fmt.Printf("%d, %d, %d", a, b, c)
}

func init(){
	c = 30
	fmt.Println(c)
}

func init(){
	a = 10
	fmt.Println(a)
}

func init(){
	b = 20
	fmt.Println(b)
}

