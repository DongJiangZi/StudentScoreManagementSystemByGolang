package main

import (
	"EduCoreStudentManagementSystem/internal/app/service"
	"EduCoreStudentManagementSystem/internal/infrastructure/config"
	"EduCoreStudentManagementSystem/internal/infrastructure/factory"
	filelogger "EduCoreStudentManagementSystem/internal/infrastructure/logger/async"
	cliHandler "EduCoreStudentManagementSystem/internal/interfaces/cli"
	"fmt"
	"os"
)

func main() {
	cfg := config.Load()

	repo, err := factory.NewStudentRepository(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "启动失败：初始化仓储失败：%v\n", err)
		os.Exit(1)
	}

	logger, err := filelogger.NewAsyncLogger(cfg.LogFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "启动失败：初始化日志失败：%v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Info("application started")
	studentService := service.NewStudentService(repo, logger)
	handler := cliHandler.NewHandler(studentService)

	handler.Run()
	logger.Info("application stopped")
}
