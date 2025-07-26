package builder

import "fmt"

type Builder interface {
	SetName(name string) Builder
	SetRollNo(rollNo int) Builder
	SetCourse(course string) Builder
	SetGradYear(gradYear int) Builder
	Build() *Student
}

type Student struct {
	name     string
	rollNo   int
	course   string
	gradYear int
}

func NewStudent() Builder {
	return &Student{}
}
func (s *Student) SetName(name string) Builder {
	s.name = name
	return s
}
func (s *Student) SetRollNo(rollNo int) Builder {
	s.rollNo = rollNo
	return s
}
func (s *Student) SetCourse(course string) Builder {
	s.course = course
	return s
}
func (s *Student) SetGradYear(gradYear int) Builder {
	s.gradYear = gradYear
	return s
}
func (s *Student) Build() *Student {
	return s
}

func BuilderDP() {
	student1 := NewStudent().
		SetName("John Doe").
		SetRollNo(12345).
		SetCourse("Computer Science").
		SetGradYear(2023).
		Build()

	fmt.Println(student1)
}
