package router

import (
	"EduCoreStudentManagementSystem/internal/interfaces/http/handler"
	response "EduCoreStudentManagementSystem/internal/interfaces/http/response"
	"net/http"
	"strings"
)

func RegisterRoutes(studentHandler *handler.StudentHandler) {
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			studentHandler.ListStudents(w, r)
		case http.MethodPost:
			studentHandler.CreateStudent(w, r)
		default:
			response.WriteJson(w, http.StatusMethodNotAllowed, response.Body{
				Code:    4050,
				Message: "method not allowed",
				Data:    nil,
			})
		}
	})

	http.HandleFunc("/students/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/students/")

		if path == "" {
			response.WriteJson(w, http.StatusBadRequest, response.Body{
				Code:    4001,
				Message: "invalid path",
				Data:    nil,
			})
			return
		}

		if strings.Contains(path, "/scores/") {
			switch r.Method {
			case http.MethodPut:
				studentHandler.UpdateStudentScore(w, r)
			default:
				response.WriteJson(w, http.StatusMethodNotAllowed, response.Body{
					Code:    4050,
					Message: "method not allowed",
					Data:    nil,
				})
			}
			return
		}

		switch r.Method {
		case http.MethodGet:
			studentHandler.GetStudent(w, r)
		case http.MethodPut:
			studentHandler.UpdateStudentInfo(w, r)
		case http.MethodDelete:
			studentHandler.DeleteStudent(w, r)
		default:
			response.WriteJson(w, http.StatusMethodNotAllowed, response.Body{
				Code:    4050,
				Message: "method not allowed",
				Data:    nil,
			})

		}
	})
}

func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
