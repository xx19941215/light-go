package model

type student struct {
	Name string
	score float64
}

func NewStudent(s string) *student {
	return &student{
		Name: s,
	}
}

//如果首字母小写，则在其他包不可以直接使用，提供一个方法供其使用
func (s *student) GetScore() float64{
	return s.score
}