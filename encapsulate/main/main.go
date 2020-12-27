package main

import (
	"fmt"
	"go.learning/m/v2/encapsulate/model"
)

func main(){
	person := model.NewPerson("Vincent")
	person.SetAge(29)
	person.SetSal(10000)

	student := model.NewStudent("Vincent")
	student.SetSal(0)
	student.SetAge(29)
	student.SetStudentNum("0112907")
	student.SetScore(100.5)
	
	fmt.Println(person)
	fmt.Println(student)
}