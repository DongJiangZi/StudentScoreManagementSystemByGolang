package test

import (
	"EduCoreStudentManagementSystem/internal/app/service"
	"EduCoreStudentManagementSystem/internal/repository/memory"
	"testing"
)

func TestCreateStudentSuccess(t *testing.T) {
	repo := memory.NewStudentRepository()
	studentService := service.NewStudentService(repo)

	student, err := studentService.CreateStudent("S001", "Alice", 20, "Female")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if student.ID != "S001" {
		t.Fatalf("expected ID S001, got %s", student.ID)
	}
}

func TestCreateStudentDuplicate(t *testing.T) {
	repo := memory.NewStudentRepository()
	studentService := service.NewStudentService(repo)

	_, err := studentService.CreateStudent("S001", "Alice", 20, "Female")
	if err != nil {
		t.Fatalf("first create should succeed, got %v", err)
	}

	_, err = studentService.CreateStudent("S001", "Bob", 22, "Male")
	if err == nil {
		t.Fatal("expected duplicate error, got nil")
	}
}

func TestUpdateStudentScoreNotFound(t *testing.T) {
	repo := memory.NewStudentRepository()
	studentService := service.NewStudentService(repo)

	err := studentService.UpdateStudentScore("S999", "Math", 90)
	if err == nil {
		t.Fatal("expected not found error, got nil")
	}
}
