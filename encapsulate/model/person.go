package model

type person struct {
	Name string
	age int
	sal float64
}

type student struct {
	person
	studentNum string
	score float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
		age: 0,
		sal: 0,
	}
}

func NewStudent(name string) *student {
	return &student{
		person: person{
			Name: "Vincent",
		},
	}
}

func (p *person) GetAge() int{
	return p.age
}

func (p *person) SetAge(age int) {
	p.age = age
}

func (p *person) GetSal() float64 {
	return p.sal
}

func (p *person) SetSal(sal float64) {
	p.sal = sal
}

func (s *student) GetStudentNum() string {
	return s.studentNum
}

func (s *student) SetStudentNum(studentNum string) {
	s.studentNum = studentNum
}

func (s *student) GetScore() float64 {
	return s.score
}

func (s *student) SetScore(score float64) {
	s.score = score
}