package main

import "fmt"

type Cat struct {
	Name string
	Age int
}

type A struct {
	Name string
}

func (a A) test(){
	fmt.Println(a.Name)
}

type B struct {
	Name string
}



func main(){
	a := A{"123"}
	fmt.Println(a)
	a.test()
	//cat := Cat{"student", 18 }
	//fmt.Println(cat)
	//
	//var a A
	//var b B
	//a = A(b)
	//fmt.Println(a)
}
