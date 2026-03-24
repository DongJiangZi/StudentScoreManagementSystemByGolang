package errs

import "errors"

var (
	ErrStudentAlreadyExists = errors.New("student already exists")
	ErrStudentNotFound      = errors.New("student not found")
	ErrInvalidScore         = errors.New("invalid score")
	ErrInvalidStudentInput  = errors.New("invalid input")
)
