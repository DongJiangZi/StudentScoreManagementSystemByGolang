package repository

import "EduCoreStudentManagementSystem/internal/domain/entity"

type StudentRepository interface {
	Create(student *entity.Student) error
	Update(student *entity.Student) error
	Delete(id string) error
	FindByID(id string) (*entity.Student, error)
	FindAll() ([]*entity.Student, error)
}
