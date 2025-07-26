package prototype

import "fmt"

// Prototype interface
type StudentPrototype interface {
	Clone() StudentPrototype
	GetDetails() string
}

// Concrete Student type
type Student struct {
	Name       string
	RollNo     int
	Course     string
	GradYear   int
	University string
}

// Clone method for Student
func (s *Student) Clone() StudentPrototype {
	// Return a deep copy of Student
	copy := *s
	return &copy
}

func (s *Student) GetDetails() string {
	return fmt.Sprintf("Name: %s, RollNo: %d, Course: %s, GradYear: %d", s.Name, s.RollNo, s.Course, s.GradYear)
}

func (s *Student) SetName(name string) {
	s.Name = name
}
func (s *Student) SetRoll(roll int) {
	s.RollNo = roll
}
func (s *Student) SetCourse(course string) {
	s.Course = course
}

func Prototype_DP() {
	// Create an original student
	studentTemplate := &Student{
		Name:       "Name",
		RollNo:     0,
		Course:     "Course Name",
		GradYear:   2024,
		University: "XYZ University",
	}

	// Clone the template and modify details
	S1 := studentTemplate.Clone().(*Student)
	S1.SetName("John Doe")
	S1.SetRoll(101)
	S1.SetCourse("Computer Science")

	// Display details of both students
	fmt.Println("Student Template:", studentTemplate.GetDetails())
	fmt.Println("Student Details:", S1.GetDetails())
}
