package main

/*
 * @Author: your name
 * @Date: 2020-11-21 20:57:41
 * @LastEditTime: 2020-12-12 21:24:05
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \go_learning\hello.go
 */

import "fmt"

// var name string
// var age int
// var isOk bool

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "dream"
	age = 16
	isOk = true

	fmt.Printf(name, age, isOk)

	a := [...]int{1: 1, 2: 3}
	fmt.Printf("%d", a)

	switch i := "y"; i {
	case "y", "Y":
		fmt.Println(i)
		fallthrough
	case "n", "N":
		fmt.Println(i)
	}

	// switch use case
	score := 85
	grade := ' '
	if score >= 90 {
		grade = 'A'
	} else if score >= 80 {
		grade = 'B'
	} else {
		grade = 'F'
	}

	fmt.Println(grade)

	for i := 1; i < 100; i++ {
		fmt.Println(i)
	}

	signal := true
	for signal {
		go func() {
			signal = false
		}()
		fmt.Println(signal)
	}

	// function test
	sum := functionTest(5, 9)
	fmt.Println(sum)

	// anonymous function calling in the
	sum = add(1, 2)
	fmt.Println(sum)

	sumTmp := add(1, 21)
	fmt.Print(sumTmp)

	// 实参到形参永远是值拷贝
	g := 10
	chvalue(g)
	fmt.Println(g)

	chpointer(&g)
	fmt.Println(g)

	var p *int = &a
	f: = *p + 1

}

func functionTest(a, b int) (sum int) {

	// sum内部函数最外层的局部变量, 由返回值声明
	sum = a + b
	// return简写, 默认返回sum
	return
}

// 匿名函数嵌套定义
func add(a, b int) (sum int) {
	anonymous := func(x, y int) int {
		return x + y
	}
	return anonymous(a, b)
}

// 实参到形参的值拷贝
func chvalue(a int) int {
	a = a + 1
	return a
}

// 实参到形参的值拷贝
func chpointer(a *int) {
	*a = *a + 1
	return
}
