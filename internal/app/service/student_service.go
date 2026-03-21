package service

import (
	"EduCoreStudentManagementSystem/internal/domain/entity"
	"EduCoreStudentManagementSystem/internal/infrastructure/logger"
	"EduCoreStudentManagementSystem/internal/repository"
	"errors"
	"sort"
	"strconv"
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

func (s *StudentService) ListStudentsByAverageDesc() ([]*entity.Student, error) {
	students, err := s.repo.FindAll()
	if err != nil {
		s.logger.Error("list students by average failed, repository error: " + err.Error())
		return nil, err
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].AverageScore() > students[j].AverageScore()
	})

	s.logger.Info("listed students by average score descending")
	return students, nil
}

func (s *StudentService) ListFailedStudents() ([]*entity.Student, error) {
	students, err := s.repo.FindAll()
	if err != nil {
		s.logger.Error("list failed students failed, repository error: " + err.Error())
		return nil, err
	}

	var failed []*entity.Student
	for _, student := range students {
		if student.HasFailedSubject() {
			failed = append(failed, student)
		}
	}

	s.logger.Info("listed failed students")
	return students, nil
}

func (s *StudentService) TopNStudentsByAverage(n int) ([]*entity.Student, error) {
	if n <= 0 {
		return []*entity.Student{}, nil
	}

	students, err := s.ListStudentsByAverageDesc()
	if err != nil {
		s.logger.Error("top N students failed: " + err.Error())
		return nil, err
	}

	if n > len(students) {
		n = len(students)
	}

	results := students[:n]
	s.logger.Info("listed top students by average, n=" + strconv.Itoa(n))
	return results, nil
}
