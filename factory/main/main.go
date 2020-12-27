package main

import (
	"fmt"
	"go.learning/m/v2/factory/model"
)

func main(){
	student := model.NewStudent("Vincent", 122)
	fmt.Println(student.Name, student.Score)
}
