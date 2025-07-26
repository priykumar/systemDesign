package prototype

import (
	"testing"
)

func TestStudentClone(t *testing.T) {
	original := &Student{
		Name:       "Template",
		RollNo:     0,
		Course:     "Course",
		GradYear:   2024,
		University: "XYZ University",
	}

	clone := original.Clone().(*Student)

	// Modify the clone
	clone.SetName("John")
	clone.SetRoll(101)
	clone.SetCourse("CS")

	// Assertions
	if clone == original {
		t.Error("Clone should be a different instance, but got the same pointer")
	}

	if original.Name == clone.Name {
		t.Error("Expected clone.Name to be different from original.Name")
	}

	if original.RollNo == clone.RollNo {
		t.Error("Expected clone.RollNo to be different from original.RollNo")
	}

	if original.Course == clone.Course {
		t.Error("Expected clone.Course to be different from original.Course")
	}

	if original.University != clone.University {
		t.Error("Expected University to be copied correctly")
	}
}
