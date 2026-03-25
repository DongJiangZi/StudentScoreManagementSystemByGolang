package memory

import (
	"EduCoreStudentManagementSystem/internal/domain/entity"
	"EduCoreStudentManagementSystem/internal/pkg/errs"
	"sync"
)

type StudentRepository struct {
	data map[string]*entity.Student
	mu   sync.RWMutex
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		data: make(map[string]*entity.Student),
	}
}

func (r *StudentRepository) Create(student *entity.Student) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[student.ID]; exists {
		return errs.ErrStudentAlreadyExists
	}

	r.data[student.ID] = student
	return nil
}

func (r *StudentRepository) Update(student *entity.Student) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[student.ID]; !exists {
		return errs.ErrStudentNotFound
	}

	r.data[student.ID] = student
	return nil
}

func (r *StudentRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; !exists {
		return errs.ErrStudentNotFound
	}

	delete(r.data, id)
	return nil
}

func (r *StudentRepository) FindByID(id string) (*entity.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	student, exists := r.data[id]
	if !exists {
		return nil, errs.ErrStudentNotFound
	}

	return student, nil
}

func (r *StudentRepository) FindAll() ([]*entity.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	students := make([]*entity.Student, 0, len(r.data))
	for _, student := range r.data {
		students = append(students, student)
	}

	return students, nil
}
