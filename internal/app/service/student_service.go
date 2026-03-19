package service

import (
	"EduCoreStudentManagementSystem/internal/domain/entity"
	"EduCoreStudentManagementSystem/internal/repository"
	"errors"
)

var (
	ErrStudentAlreadyExists = errors.New("student already exists")
	ErrStudentNotFound      = errors.New("student not found")
)

type StudentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) *StudentService {
	return &StudentService{
		repo: repo,
	}
}

func (s *StudentService) CreateStudent(id, name string, age int, gender string) (*entity.Student, error) {
	existing, err := s.repo.FindByID(id)
	if err == nil && existing != nil {
		return nil, ErrStudentAlreadyExists
	}

	student, err := entity.NewStudent(id, name, age, gender)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(student); err != nil {
		return nil, err
	}

	return student, nil
}

func (s *StudentService) GetStudent(id string) (*entity.Student, error) {
	student, err := s.repo.FindByID(id)
	if err != nil {
		return nil, ErrStudentNotFound
	}

	return student, nil
}

func (s *StudentService) ListStudents() ([]*entity.Student, error) {
	return s.repo.FindAll()
}

func (s *StudentService) UpdateStudentScore(id, subject string, score float64) error {
	student, err := s.repo.FindByID(id)
	if err != nil {
		return ErrStudentNotFound
	}

	if err := student.UpdateScore(subject, score); err != nil {
		return err
	}

	if err := s.repo.Update(student); err != nil {
		return nil
	}

	return nil
}

func (s *StudentService) DeleteStudent(id string) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}
