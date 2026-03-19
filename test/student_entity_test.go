package test

import (
	"EduCoreStudentManagementSystem/internal/domain/entity"
	"testing"
)

func TestNewStudentSuccess(t *testing.T) {
	student, err := entity.NewStudent("S001", "Alice", 20, "Female")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if student.ID != "S001" {
		t.Fatalf("expected ID S001, got %s", student.ID)
	}

	if student.Name != "Alice" {
		t.Fatalf("expected Name Alice, got %s", student.Name)
	}

	if student.Age != 20 {
		t.Fatalf("expected Age 20, got %d", student.Age)
	}

	if student.Gender != "Female" {
		t.Fatalf("expected Gender Female, got %s", student.Gender)
	}
}

func TestNewStudentInvalidName(t *testing.T) {
	_, err := entity.NewStudent("S001", "", 20, "Female")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestUpdateScoreSuccess(t *testing.T) {
	student, err := entity.NewStudent("S001", "Alice", 20, "Female")
	if err != nil {
		t.Fatalf("create student failed: %v", err)
	}

	err = student.UpdateScore("Math", 95)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	score, exists := student.Scores["Math"]
	if !exists {
		t.Fatal("expected Math score to exist")
	}

	if score != 95 {
		t.Fatalf("expected Math score 95, got %v", score)
	}
}

func TestAverageScore(t *testing.T) {
	student, err := entity.NewStudent("S001", "Alice", 20, "Female")
	if err != nil {
		t.Fatalf("create student failed: %v", err)
	}

	_ = student.UpdateScore("Math", 100)
	_ = student.UpdateScore("English", 80)

	avg := student.AverageScore()
	expected := 90.0

	if avg != expected {
		t.Fatalf("expected average %.2f, got %.2f", expected, avg)
	}
}
