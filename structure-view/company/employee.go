package company

import (
	"fmt"
	"regexp"
)

// todo open the 'Structure' tool window to have a
type Employee struct {
	Name   string
	Age    int
	Salary float64
	Email  string
}

func (e *Employee) Work() {
	fmt.Printf("%s is working.\n", e.Name)
}

func (e *Employee) GetDetails() string {
	return fmt.Sprintf("Employee: %s, Age: %d, Salary: %.2f, Email: %s", e.Name, e.Age, e.Salary, e.Email)
}

func (e *Employee) IsValidEmail() bool {
	// todo using AI Assistant explain the regular expression
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(e.Email)
}
