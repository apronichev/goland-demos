package company

import "fmt"

type Manager struct {
	Employee
	Department string
}

func (m *Manager) Work() {
	fmt.Printf("%s is managing the %s department.\n", m.Name, m.Department)
}

func (m *Manager) GetDetails() string {
	return fmt.Sprintf("Manager: %s, Age: %d, Salary: %.2f, Email: %s, Department: %s", m.Name, m.Age, m.Salary, m.Email, m.Department)
}
