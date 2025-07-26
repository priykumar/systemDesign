package builder

import (
	"testing"
)

func TestBuilderDP(t *testing.T) {
	expected := &Student{
		name:     "John Doe",
		rollNo:   12345,
		course:   "Computer Science",
		gradYear: 2023,
	}

	student := NewStudent().
		SetName("John Doe").
		SetRollNo(12345).
		SetCourse("Computer Science").
		SetGradYear(2023).
		Build()

	if student.name != expected.name {
		t.Errorf("Expected name %s, got %s", expected.name, student.name)
	}
	if student.rollNo != expected.rollNo {
		t.Errorf("Expected rollNo %d, got %d", expected.rollNo, student.rollNo)
	}
	if student.course != expected.course {
		t.Errorf("Expected course %s, got %s", expected.course, student.course)
	}
	if student.gradYear != expected.gradYear {
		t.Errorf("Expected gradYear %d, got %d", expected.gradYear, student.gradYear)
	}
}
