package main

import (
	"EduCoreStudentManagementSystem/internal/app/service"
	"EduCoreStudentManagementSystem/internal/infrastructure/config"
	"EduCoreStudentManagementSystem/internal/infrastructure/factory"
	filelogger "EduCoreStudentManagementSystem/internal/infrastructure/logger/async"
	"EduCoreStudentManagementSystem/internal/interfaces/http/handler"
	"EduCoreStudentManagementSystem/internal/interfaces/http/router"
	"fmt"
	"net/http"
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

	logger.Info("http server starting...")

	studentService := service.NewStudentService(repo, logger)

	studentHandler := handler.NewStudentHandler(studentService)

	router.RegisterRoutes(studentHandler)

	addr := ":8080"
	fmt.Println("HTTP server listening on", addr)

	mux := http.DefaultServeMux
	if err := http.ListenAndServe(addr, router.WithCORS(mux)); err != nil {
		fmt.Fprintf(os.Stderr, "HTTP 服务启动失败：%v\n", err)
		os.Exit(1)
	}
}
