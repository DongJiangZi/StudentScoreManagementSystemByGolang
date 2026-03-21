package factory

import (
	"EduCoreStudentManagementSystem/internal/infrastructure/config"
	"EduCoreStudentManagementSystem/internal/repository"
	"EduCoreStudentManagementSystem/internal/repository/json"
	"EduCoreStudentManagementSystem/internal/repository/memory"
	"errors"
)

func NewStudentRepository(cfg *config.Config) (repository.StudentRepository, error) {
	switch cfg.RepositoryType {
	case "memory":
		return memory.NewStudentRepository(), nil
	case "json":
		return json.NewStudentRepository(cfg.DataFilePath)
	default:
		return nil, errors.New("unsupported repository type: " + cfg.RepositoryType)
	}
}
