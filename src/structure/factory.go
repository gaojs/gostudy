// 工厂模式
package structure

import "fmt"

// 学生（姓名，分数）
type Student struct {
	name  string
	score int
}

func NewStudent(name string, score int) *Student {
	return &Student{name, score}
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) GetScore() int {
	return s.score
}

func (s *Student) SetScore(score int) {
	s.score = score
}

func (s *Student) Show() {
	fmt.Printf("Student: %T\n", s)
	fmt.Println(s.name, s.score)
}

// 小学生
type Pupil struct {
	Student // 嵌入匿名结构体
}

func (s *Pupil) Show() {
	fmt.Printf("Pupil: %T\n", s)
	s.Student.Show()
}

// 大学生
type Graduate struct {
	Student // 嵌入匿名结构体
}

func (s *Graduate) Show() {
	fmt.Printf("Graduate: %T\n", s)
	s.Student.Show()
}
