package entity

import (
	"errors"
	"strings"
	"time"
)

type Student struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Age       int                `json:"age"`
	Gender    string             `json:"gender"`
	Scores    map[string]float64 `json:"scores"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

func NewStudent(id, name string, age int, gender string) (*Student, error) {
	student := &Student{
		ID:        strings.TrimSpace(id),
		Name:      strings.TrimSpace(name),
		Age:       age,
		Gender:    strings.TrimSpace(gender),
		Scores:    make(map[string]float64),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := student.Validate(); err != nil {
		return nil, err
	}

	return student, nil
}

func (s *Student) Validate() error {
	if strings.TrimSpace(s.ID) == "" {
		return errors.New("student ID cannot be empty")
	}

	if strings.TrimSpace(s.Name) == "" {
		return errors.New("student name cannot be empty")
	}

	if s.Age <= 0 {
		return errors.New("student age must be greater than 0")
	}

	if strings.TrimSpace(s.Gender) == "" {
		return errors.New("student gender cannot be empty")
	}

	return nil
}

func (s *Student) UpdateScore(subject string, score float64) error {
	subject = strings.TrimSpace(subject)
	if subject == "" {
		return errors.New("subject cannot be empty")
	}

	if score < 0 || score > 100 {
		return errors.New("score must be between 0 and 100")
	}

	if s.Scores == nil {
		s.Scores = make(map[string]float64)
	}

	s.Scores[subject] = score
	s.UpdatedAt = time.Now()
	return nil
}

func (s *Student) AverageScore() float64 {
	if len(s.Scores) == 0 {
		return 0
	}

	var total float64
	for _, score := range s.Scores {
		total += score
	}

	return total / float64(len(s.Scores))
}

func (s *Student) UpdateBasicInfo(name string, age int, gender string) error {
	name = strings.TrimSpace(name)
	gender = strings.TrimSpace(gender)

	if name == "" {
		return errors.New("name cannot be empty")
	}

	if age <= 0 {
		return errors.New("age must be greater than 0")
	}

	if gender == "" {
		return errors.New("gender cannot be empty")
	}

	s.Name = name
	s.Age = age
	s.Gender = gender
	s.UpdatedAt = time.Now()

	return nil
}

func (s *Student) HasFailedSubject() bool {
	for _, score := range s.Scores {
		if score < 60 {
			return true
		}
	}
	return false
}

func (s *Student) HasScores() bool {
	return len(s.Scores) > 0
}

func (s *Student) GetScore(subject string) (float64, bool) {
	score, ok := s.Scores[subject]
	return score, ok
}
