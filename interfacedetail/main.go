package main

import "fmt"

type Student interface {
	Study()
	DoHomework()
}

type Pupil struct {

}

func (this *Pupil) Study(){
	fmt.Println("pupil studying")
}

func (this *Pupil) DoHomework(){
	fmt.Println("pupil doing homework")
}

type Graduate struct {

}

func (this *Graduate) Study(){
	fmt.Println("graduate studying")
}

func (this *Graduate) DoHomework(){
	fmt.Println("graduate doing homework")
}

func doSomething(stu Student){
	stu.Study()
	stu.DoHomework()
}

func main(){
	pupil := &Pupil{}
	doSomething(pupil)
	graduate := &Graduate{}
	doSomething(graduate)

	var a interface {}
	a = pupil
	var b *Pupil
	b = a.(*Pupil)
	b.DoHomework()
	b.Study()
}
