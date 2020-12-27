package main

import "fmt"

var DEFAULT_STUDENT_NAME = "Vincent"

const (
	A = iota
	B
	C
	D
)

func main()	{
	r := []rune("Hello, 成都")
	for index := range r{
		fmt.Printf("%c\n", r[index])
	}
}

func TestDup1()	{
	fmt.Println("")
}
