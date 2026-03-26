package service

type CreateStudentRequest struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

type UpdateStudentInfoRequest struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

type UpdateStudentScoreRequest struct {
	ID      string  `json:"id"`
	Subject string  `json:"subject"`
	Score   float64 `json:"score"`
}
