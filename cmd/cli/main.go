package main

import (
	"EduCoreStudentManagementSystem/internal/app/service"
	"EduCoreStudentManagementSystem/internal/infrastructure/config"
	filelogger "EduCoreStudentManagementSystem/internal/infrastructure/logger/async"
	cliHandler "EduCoreStudentManagementSystem/internal/interfaces/cli"
	jsonrepo "EduCoreStudentManagementSystem/internal/repository/json"
)

func main() {
	cfg := config.Load()

	repo, err := jsonrepo.NewStudentRepository(cfg.DataFilePath)
	if err != nil {
		panic(err)
	}

	logger, err := filelogger.NewAsyncLogger(cfg.LogFilePath)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	studentService := service.NewStudentService(repo, logger)
	handler := cliHandler.NewHandler(studentService)

	handler.Run()
}
