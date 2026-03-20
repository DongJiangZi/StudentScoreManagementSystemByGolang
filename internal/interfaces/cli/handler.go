package cli

import (
	"EduCoreStudentManagementSystem/internal/app/service"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Handler struct {
	studentService *service.StudentService
	reader         *bufio.Reader
}

func NewHandler(studentService *service.StudentService) *Handler {
	return &Handler{
		studentService: studentService,
		reader:         bufio.NewReader(os.Stdin),
	}
}

func (h *Handler) Run() {
	for {
		h.printMenu()
		choice := h.readLine("Please choose: ")

		switch choice {
		case "1":
			h.handleCreateStudent()
		case "2":
			h.handleGetStudent()
		case "3":
			h.handleListStudents()
		case "4":
			h.handleUpdateScore()
		case "5":
			h.handleDeleteStudent()
		case "6":
			h.handleUpdateStudentInfo()
		case "0":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}

		fmt.Println()
	}
}

func (h *Handler) printMenu() {
	fmt.Println("====== Student Management System ======")
	fmt.Println("1. Create student")
	fmt.Println("2. Get student by ID")
	fmt.Println("3. List all students")
	fmt.Println("4. Update student score")
	fmt.Println("5. Delete student")
	fmt.Println("6. Update student info")
	fmt.Println("0. Exit")
}

func (h *Handler) handleCreateStudent() {
	id := h.readLine("Enter student ID: ")
	name := h.readLine("Enter student name: ")
	ageText := h.readLine("Enter student age: ")
	gender := h.readLine("Enter student gender: ")

	age, err := strconv.Atoi(ageText)
	if err != nil {
		fmt.Println("Invalid age.")
		return
	}

	student, err := h.studentService.CreateStudent(id, name, age, gender)
	if err != nil {
		fmt.Println("Create student failed:", err)
		return
	}

	fmt.Println("Student created successfully:")
	fmt.Printf("ID: %s, Name: %s, Age: %d, Gender: %s\n",
		student.ID, student.Name, student.Age, student.Gender)
}

func (h *Handler) handleGetStudent() {
	id := h.readLine("Enter student ID: ")

	student, err := h.studentService.GetStudent(id)
	if err != nil {
		fmt.Println("Get Student failed: ", err)
	}

	fmt.Println("Student found:")
	fmt.Printf("ID: %s\n", student.ID)
	fmt.Printf("Name: %s\n", student.Name)
	fmt.Printf("Age: %d\n", student.Age)
	fmt.Printf("Gender: %s\n", student.Gender)
	fmt.Printf("Scores: %+v\n", student.Scores)
	fmt.Printf("Average Score: %.2f\n", student.AverageScore())
}

func (h *Handler) handleListStudents() {
	students, err := h.studentService.ListStudents()
	if err != nil {
		fmt.Println("List students failed: ", err)
		return
	}

	if len(students) == 0 {
		fmt.Println("No students found.")
		return
	}

	fmt.Println("All students:")
	for _, student := range students {
		fmt.Printf("ID: %s, Name: %s, Age: %d, Gender: %s, Avg: %.2f\n",
			student.ID, student.Name, student.Age, student.Gender, student.AverageScore())
	}
}

func (h *Handler) handleUpdateScore() {
	id := h.readLine("Enter student ID: ")
	subject := h.readLine("Enter subject: ")
	scoreText := h.readLine("Enter score: ")

	score, err := strconv.ParseFloat(scoreText, 64)
	if err != nil {
		fmt.Println("Invalid score.")
		return
	}

	err = h.studentService.UpdateStudentScore(id, subject, score)
	if err != nil {
		fmt.Println("Update score failed:", err)
		return
	}

	fmt.Println("Score updated successfully.")
}

func (h *Handler) handleDeleteStudent() {
	id := h.readLine("Enter student ID: ")

	err := h.studentService.DeleteStudent(id)
	if err != nil {
		fmt.Println("Delete student failed:", err)
		return
	}

	fmt.Println("Student deleted successfully.")
}

func (h *Handler) handleUpdateStudentInfo() {
	id := h.readLine("Enter student ID: ")
	name := h.readLine("Enter new name: ")
	ageText := h.readLine("Enter new age: ")
	gender := h.readLine("Enter new gender: ")

	age, err := strconv.Atoi(ageText)
	if err != nil {
		fmt.Println("Invalid age.")
		return
	}

	err = h.studentService.UpdateStudentInfo(id, name, age, gender)
	if err != nil {
		fmt.Println("Update student info failed:", err)
		return
	}

	fmt.Println("Student info updated successfully.")
}

func (h *Handler) readLine(prompt string) string {
	fmt.Print(prompt)
	input, _ := h.reader.ReadString('\n')
	return strings.TrimSpace(input)
}
