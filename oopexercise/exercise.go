package main

import "fmt"

type Student struct {
	Name string
	Age int
	Address string
	Intelligence int
}

func (student * Student ) introduction () {
	fmt.Printf("I'am %v %v %v %v",
		student.Name,
		student.Age,
		student.Address,
		student.Intelligence,
		)
}

type Visitor struct {
	Name string
	Age int
}

func (visitor *Visitor) showPrice()	{
	if visitor.Age >= 90 || visitor.Age <= 8 {
		fmt.Printf(" age %v, age  %v, Not suggested\n", visitor.Name, visitor.Age)
		return
	}
	if visitor.Age > 18 {
		fmt.Printf(" age %v, age  %v, Price 20$\n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf(" age %v, age  %v, Price free\n", visitor.Name, visitor.Age)
	}
}

func main()	{
	var v Visitor
	for ;; {
		fmt.Println("enter your name")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("exit")
			break
		}
		fmt.Println("enter your age")
		fmt.Scanln(&v.Age)
		v.showPrice()


	}
	//student := Student{
	//	Name: "Vincent",
	//	Age: 29,
	//	Address: "Chengdu, Sichuan",
	//	Intelligence: 145,
	//}
	//student.introduction()
}
