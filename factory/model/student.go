package model


type student struct {
	Name string
	Score float64
}

// 工厂模式

func NewStudent(n string, s float64) *student {
	return &student{
		Name: n,
		Score: s,
	}
}