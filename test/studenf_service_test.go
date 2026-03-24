package test

import (
	service2 "EduCoreStudentManagementSystem/internal/app/service"
	"EduCoreStudentManagementSystem/internal/repository/memory"
	"fmt"
	"testing"
)

type NoopLogger struct{}

func NewNoopLogger() *NoopLogger {
	return &NoopLogger{}
}

func (l *NoopLogger) Info(msg string) {}
func (l *NoopLogger) Error(msg string) {

}

func TestCreateStudentSuccess(t *testing.T) {
	repo := memory.NewStudentRepository()
	logger := NewNoopLogger()
	service := service2.NewStudentService(repo, logger)

	req := service2.CreateStudentRequest{
		ID:     "S001",
		Name:   "Alice",
		Age:    20,
		Gender: "Female",
	}

	student, err := service.CreateStudent(req)
	if err != nil {
		t.Fatal("expected no error, got %v", err)
	}

	if student.ID != "S001" {
		t.Fatal("expected S001, got %v", student.ID)
	}
}

func TestCreateStudentDuplicate(t *testing.T) {
	repo := memory.NewStudentRepository()
	logger := NewNoopLogger()
	service := service2.NewStudentService(repo, logger)

	req := service2.CreateStudentRequest{
		ID:     "S001",
		Name:   "Alice",
		Age:    20,
		Gender: "Female",
	}

	_, _ = service.CreateStudent(req)
	_, err := service.CreateStudent(req)
	if err == nil {
		t.Fatal("expected duplicate error, got nil")
	}
}

func TestUpdateStudentScore(t *testing.T) {
	repo := memory.NewStudentRepository()
	logger := NewNoopLogger()
	service := service2.NewStudentService(repo, logger)

	createReq := service2.CreateStudentRequest{
		ID:     "S001",
		Name:   "Alice",
		Age:    20,
		Gender: "Femail",
	}

	_, _ = service.CreateStudent(createReq)

	updateReq := service2.UpdateStudentScoreRequest{
		ID:      "S001",
		Subject: "Math",
		Score:   95,
	}

	err := service.UpdateStudentScore(updateReq)
	if err != nil {
		t.Fatal("expected no error, got %v", err)
	}
}

func TestTopNStudents(t *testing.T) {
	repo := memory.NewStudentRepository()
	logger := NewNoopLogger()
	service := service2.NewStudentService(repo, logger)

	for i := 1; i <= 3; i++ {
		id := fmt.Sprintf("S00%d", i)
		req := service2.CreateStudentRequest{
			ID:     id,
			Name:   "Student",
			Age:    20,
			Gender: "Male",
		}
		s, _ := service.CreateStudent(req)
		s.UpdateScore("Math", float64(80+i))
		repo.Update(s)
	}

	result, err := service.TopNStudentsByAverage(2)
	if err != nil {
		t.Fatal("expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Fatal("expected 2 students, got %d", len(result))
	}
}
