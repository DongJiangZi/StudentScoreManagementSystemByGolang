package service

import (
	"EduCoreStudentManagementSystem/internal/domain/entity"
	"EduCoreStudentManagementSystem/internal/infrastructure/logger"
	"EduCoreStudentManagementSystem/internal/pkg/errs"
	"EduCoreStudentManagementSystem/internal/repository"
	"sort"
	"strconv"
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

func (s *StudentService) CreateStudent(req CreateStudentRequest) (*entity.Student, error) {
	s.logger.Info("creating student: " + req.ID)

	existing, err := s.repo.FindByID(req.ID)
	if err == nil && existing != nil {
		return nil, errs.ErrStudentAlreadyExists
	}

	student, err := entity.NewStudent(req.ID, req.Name, req.Age, req.Gender)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(student); err != nil {
		s.logger.Error("create student failed: " + err.Error())
		return nil, err
	}

	s.logger.Info("student created: " + req.ID)
	return student, nil
}

func (s *StudentService) GetStudent(id string) (*entity.Student, error) {
	student, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errs.ErrStudentNotFound
	}

	return student, nil
}

func (s *StudentService) ListStudents() ([]*entity.Student, error) {
	return s.repo.FindAll()
}

func (s *StudentService) UpdateStudentScore(req UpdateStudentScoreRequest) error {
	student, err := s.repo.FindByID(req.ID)
	if err != nil {
		return errs.ErrStudentNotFound
	}

	if err := student.UpdateScore(req.Subject, req.Score); err != nil {
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

func (s *StudentService) UpdateStudentInfo(req UpdateStudentInfoRequest) error {
	student, err := s.repo.FindByID(req.ID)
	if err != nil {
		return errs.ErrStudentNotFound
	}

	if err := student.UpdateBasicInfo(req.Name, req.Age, req.Gender); err != nil {
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
