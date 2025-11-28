package company

import (
	"fmt"
	"regexp"
)

// todo open the 'Structure' tool window
type Employee struct {
	NAME   string
	Age    int
	Salary float64
	Email  string
}

type Manager struct {
	Employee
	Department string
}

type Documents struct {
	Passport    string
	SocSecurity string
	IDCard      string
	BirthCert   string
}

type Department struct {
	Name      string
	Manager   *Manager
	Employees []Employee
	Location  string
}

func (e *Employee) Work() {
	fmt.Printf("%s is working.\n", e.NAME)
}

func (e *Employee) GetDetails() string {
	return fmt.Sprintf("Employee: %s, Age: %d, Salary: %.2f, Email: %s", e.NAME, e.Age, e.Salary, e.Email)
}

func (e *Employee) IsValidEmail() bool {
	// todo using AI Assistant explain the regular expression
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(e.Email)
}
