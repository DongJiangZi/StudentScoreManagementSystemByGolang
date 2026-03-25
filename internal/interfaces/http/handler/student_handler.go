package handler

import (
	"EduCoreStudentManagementSystem/internal/app/service"
	response "EduCoreStudentManagementSystem/internal/interfaces/http/response"
	"net/http"
)

type StudentHandler struct {
	studentService *service.StudentService
}

func NewStudentHandler(studentService *service.StudentService) *StudentHandler {
	return &StudentHandler{
		studentService: studentService,
	}
}

func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodGet {
		response.WriteJson(w, http.StatusMethodNotAllowed, response.Body{
			Code: 4050,
			Message: "method not allowed",
			Data: nil,
		})
		return
	}
}