package response

import (
	"EduCoreStudentManagementSystem/internal/pkg/errs"
	"encoding/json"
	"errors"
	"net/http"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WriteJson(w http.ResponseWriter, httpStatus int, body Body) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatus)

	_ = json.NewEncoder(w).Encode(body)
}

func Success(w http.ResponseWriter, data interface{}) {
	WriteJson(w, http.StatusOK, Body{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func Created(w http.ResponseWriter, data interface{}) {
	WriteJson(w, http.StatusCreated, Body{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func Error(w http.ResponseWriter, err error) {
	httpStatus, code, message := mapError(err)

	WriteJson(w, httpStatus, Body{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func mapError(err error) (int, int, string) {
	switch {
	case errors.Is(err, errs.ErrStudentNotFound):
		return http.StatusNotFound, 4041, "student not found"
	case errors.Is(err, errs.ErrStudentAlreadyExists):
		return http.StatusConflict, 4091, "student already exists"
	case errors.Is(err, errs.ErrInvalidStudentInput):
		return http.StatusBadRequest, 4001, "invalid student input"
	case errors.Is(err, errs.ErrInvalidScore):
		return http.StatusBadRequest, 4002, "invalid score"
	default:
		return http.StatusInternalServerError, 5001, "internal server error"
	}
}
