package service

import (
	"EduCoreStudentManagementSystem/internal/domain/entity"
	"EduCoreStudentManagementSystem/internal/infrastructure/logger"
	"EduCoreStudentManagementSystem/internal/pkg/errs"
	"EduCoreStudentManagementSystem/internal/repository"
	"errors"
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
	s.logger.Info("creating student, id=" + req.ID + ", name=" + req.Name)

	existing, err := s.repo.FindByID(req.ID)
	if err == nil && existing != nil {
		s.logger.Warn("create student rejected, duplicate id=" + req.ID)
		return nil, errs.ErrStudentAlreadyExists
	}
	if err != nil && !errors.Is(err, errs.ErrStudentNotFound) {
		s.logger.Error("create student failed while checking existing student, id=" + req.ID + ", err=" + err.Error())
		return nil, err
	}

	student, err := entity.NewStudent(req.ID, req.Name, req.Age, req.Gender)
	if err != nil {
		s.logger.Warn("create student validation failed, id=" + req.ID + ", err=" + err.Error())
		return nil, err
	}

	if err := s.repo.Create(student); err != nil {
		s.logger.Error("create student failed while saving, id=" + req.ID + ", err=" + err.Error())
		return nil, err
	}

	s.logger.Info("student created, id=" + req.ID)
	return student, nil
}

func (s *StudentService) GetStudent(id string) (*entity.Student, error) {
	s.logger.Info("getting student, id=" + id)
	student, err := s.repo.FindByID(id)
	if err != nil {
		s.logger.Warn("get student not found, id=" + id + ", err=" + err.Error())
		return nil, errs.ErrStudentNotFound
	}

	s.logger.Info("student fetched, id=" + id)
	return student, nil
}

func (s *StudentService) ListStudents() ([]*entity.Student, error) {
	students, err := s.repo.FindAll()
	if err != nil {
		s.logger.Error("list students failed, err=" + err.Error())
		return nil, err
	}

	s.logger.Info("students listed, count=" + strconv.Itoa(len(students)))
	return students, nil
}

func (s *StudentService) UpdateStudentScore(req UpdateStudentScoreRequest) error {
	s.logger.Info("updating score, id=" + req.ID + ", subject=" + req.Subject)
	student, err := s.repo.FindByID(req.ID)
	if err != nil {
		s.logger.Warn("update score rejected, student not found, id=" + req.ID)
		return errs.ErrStudentNotFound
	}

	if err := student.UpdateScore(req.Subject, req.Score); err != nil {
		s.logger.Warn("update score validation failed, id=" + req.ID + ", err=" + err.Error())
		return err
	}

	if err := s.repo.Update(student); err != nil {
		s.logger.Error("update score failed while saving, id=" + req.ID + ", err=" + err.Error())
		return err
	}

	s.logger.Info("score updated, id=" + req.ID + ", subject=" + req.Subject)
	return nil
}

func (s *StudentService) DeleteStudent(id string) error {
	s.logger.Info("deleting student, id=" + id)
	_, err := s.repo.FindByID(id)
	if err != nil {
		s.logger.Warn("delete rejected, student not found, id=" + id)
		return err
	}

	if err := s.repo.Delete(id); err != nil {
		s.logger.Error("delete student failed, id=" + id + ", err=" + err.Error())
		return err
	}

	s.logger.Info("student deleted, id=" + id)
	return nil
}

func (s *StudentService) UpdateStudentInfo(req UpdateStudentInfoRequest) error {
	s.logger.Info("updating student info, id=" + req.ID)
	student, err := s.repo.FindByID(req.ID)
	if err != nil {
		s.logger.Warn("update student info rejected, student not found, id=" + req.ID)
		return errs.ErrStudentNotFound
	}

	if err := student.UpdateBasicInfo(req.Name, req.Age, req.Gender); err != nil {
		s.logger.Warn("update student info validation failed, id=" + req.ID + ", err=" + err.Error())
		return err
	}

	if err := s.repo.Update(student); err != nil {
		s.logger.Error("update student info failed while saving, id=" + req.ID + ", err=" + err.Error())
		return err
	}

	s.logger.Info("student info updated, id=" + req.ID)
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

	s.logger.Info("listed failed students, count=" + strconv.Itoa(len(failed)))
	return failed, nil
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
