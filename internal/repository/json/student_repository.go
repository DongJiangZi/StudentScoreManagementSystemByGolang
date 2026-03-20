package json

import (
	"EduCoreStudentManagementSystem/internal/domain/entity"
	"encoding/json"
	"errors"
	"os"
	"sync"
)

var (
	ErrStudentAlreadyExists = errors.New("student already exists")
	ErrStudentNotFound      = errors.New("student not found")
)

type StudentRepository struct {
	filePath string
	data     map[string]*entity.Student
	mu       sync.RWMutex
}

func NewStudentRepository(filePath string) (*StudentRepository, error) {
	repo := &StudentRepository{
		filePath: filePath,
		data:     make(map[string]*entity.Student),
	}

	if err := repo.load(); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *StudentRepository) load() error {
	file, err := os.Open(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	var students []*entity.Student
	if err := json.NewDecoder(file).Decode(&students); err != nil {
		return err
	}

	for _, s := range students {
		r.data[s.ID] = s
	}

	return nil
}

func (r *StudentRepository) save() error {
	students := make([]*entity.Student, 0, len(r.data))
	for _, s := range r.data {
		students = append(students, s)
	}

	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(students)
}

func (r *StudentRepository) Create(student *entity.Student) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[student.ID]; exists {
		return ErrStudentAlreadyExists
	}

	r.data[student.ID] = student
	return r.save()
}

func (r *StudentRepository) Update(student *entity.Student) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[student.ID]; !exists {
		return ErrStudentNotFound
	}

	r.data[student.ID] = student
	return r.save()
}

func (r *StudentRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; !exists {
		return ErrStudentNotFound
	}

	delete(r.data, id)
	return r.save()
}

func (r *StudentRepository) FindByID(id string) (*entity.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	student, exists := r.data[id]
	if !exists {
		return nil, ErrStudentNotFound
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
