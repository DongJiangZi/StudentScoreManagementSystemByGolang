package main

import (
	"EduCoreStudentManagementSystem/internal/app/service"
	cliHandler "EduCoreStudentManagementSystem/internal/interfaces/cli"
	"EduCoreStudentManagementSystem/internal/repository/memory"
)

func main() {
	repo := memory.NewStudentRepository()
	studentService := service.NewStudentService(repo)
	handler := cliHandler.NewHandler(studentService)

	handler.Run()
}
