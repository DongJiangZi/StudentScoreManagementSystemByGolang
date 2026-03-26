package handler

import (
	"EduCoreStudentManagementSystem/internal/app/service"
	response "EduCoreStudentManagementSystem/internal/interfaces/http/response"
	"encoding/json"
	"net/http"
	"strings"
)

type StudentHandler struct {
	studentService *service.StudentService
}

func NewStudentHandler(studentService *service.StudentService) *StudentHandler {
	return &StudentHandler{
		studentService: studentService,
	}
}

func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteJson(w, http.StatusMethodNotAllowed, response.Body{
			Code:    4050,
			Message: "method not allowed",
			Data:    nil,
		})
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/students/")
	if id == "" {
		response.WriteJson(w, http.StatusBadRequest, response.Body{
			Code:    4001,
			Message: "student id is required",
			Data:    nil,
		})
		return
	}

	student, err := h.studentService.GetStudent(id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, student)
}

func (h *StudentHandler) ListStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteJson(w, http.StatusMethodNotAllowed, response.Body{
			Code:    4050,
			Message: "method not allowed",
			Data:    nil,
		})
		return
	}

	students, _ := h.studentService.ListStudents()
	response.Success(w, students)
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteJson(w, http.StatusMethodNotAllowed, response.Body{
			Code:    4050,
			Message: "method not allowed",
			Data:    nil,
		})
		return
	}

	var req service.CreateStudentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.Body{
			Code:    4002,
			Message: "invalid request body",
			Data:    nil,
		})
	}

	student, err := h.studentService.CreateStudent(req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Created(w, student)
}

func (h *StudentHandler) UpdateStudentInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response.WriteJson(w, http.StatusMethodNotAllowed, response.Body{
			Code:    4050,
			Message: "method not allowed",
			Data:    nil,
		})
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/students/")
	if id == "" {
		response.WriteJson(w, http.StatusBadRequest, response.Body{
			Code:    4001,
			Message: "student id is required",
			Data:    nil,
		})
		return
	}

	var req service.UpdateStudentInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.Body{
			Code:    4002,
			Message: "invalid request body",
			Data:    nil,
		})
		return
	}

	req.ID = id

	err := h.studentService.UpdateStudentInfo(req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.WriteJson(w, http.StatusMethodNotAllowed, response.Body{
			Code:    4050,
			Message: "method not allowed",
			Data:    nil,
		})
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/students/")
	if id == "" {
		response.WriteJson(w, http.StatusBadRequest, response.Body{
			Code:    4001,
			Message: "student id is required",
			Data:    nil,
		})
		return
	}

	if err := h.studentService.DeleteStudent(id); err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *StudentHandler) UpdateStudentScore(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response.WriteJson(w, http.StatusMethodNotAllowed, response.Body{
			Code:    4050,
			Message: "method not allowed",
			Data:    nil,
		})
		return
	}

	// 期望路径: /students/{id}/scores/{subject}
	path := strings.TrimPrefix(r.URL.Path, "/students/")
	parts := strings.Split(path, "/")
	if len(parts) != 3 || parts[1] != "scores" || parts[0] == "" || parts[2] == "" {
		response.WriteJson(w, http.StatusBadRequest, response.Body{
			Code:    4001,
			Message: "invalid score update path",
			Data:    nil,
		})
		return
	}

	id := parts[0]
	subject := parts[2]

	var req service.UpdateStudentScoreRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.Body{
			Code:    4002,
			Message: "invalid request body",
			Data:    nil,
		})
		return
	}

	req.ID = id
	req.Subject = subject

	err := h.studentService.UpdateStudentScore(req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}
