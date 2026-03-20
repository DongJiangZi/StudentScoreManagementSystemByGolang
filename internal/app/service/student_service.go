package service

import (
	"EduCoreStudentManagementSystem/internal/domain/entity"
	"EduCoreStudentManagementSystem/internal/infrastructure/logger"
	"EduCoreStudentManagementSystem/internal/repository"
	"errors"
)

var (
	ErrStudentAlreadyExists = errors.New("student already exists")
	ErrStudentNotFound      = errors.New("student not found")
)

type StudentService struct {
	repo   repository.StudentRepository
	logger logger.Logger
}

func NewStudentService(repo repository.StudentRepository, logger logger.Logger) *StudentService {
	return &StudentService{
		repo:   repo,
		logger: logger,
	}
}

func (s *StudentService) CreateStudent(id, name string, age int, gender string) (*entity.Student, error) {
	s.logger.Info("creating student: " + id)

	existing, err := s.repo.FindByID(id)
	if err == nil && existing != nil {
		return nil, ErrStudentAlreadyExists
	}

	student, err := entity.NewStudent(id, name, age, gender)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(student); err != nil {
		s.logger.Error("create student failed: " + err.Error())
		return nil, err
	}

	s.logger.Info("student created: " + id)
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

func (s *StudentService) UpdateStudentInfo(id, name string, age int, gender string) error {
	student, err := s.repo.FindByID(id)
	if err != nil {
		return ErrStudentNotFound
	}

	if err := student.UpdateBasicInfo(name, age, gender); err != nil {
		return err
	}

	if err := s.repo.Update(student); err != nil {
		return err
	}

	return nil
}
