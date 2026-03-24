package service

type CreateStudentRequest struct {
	ID     string
	Name   string
	Age    int
	Gender string
}

type UpdateStudentInfoRequest struct {
	ID     string
	Name   string
	Age    int
	Gender string
}

type UpdateStudentScoreRequest struct {
	ID      string
	Subject string
	Score   float64
}
