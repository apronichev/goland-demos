package company

import "fmt"

type Department struct {
	Name      string
	Manager   *Manager
	Employees []Employee
}

func (d *Department) AddEmployee(e Employee) {
	d.Employees = append(d.Employees, e)
}

func (d *Department) GetDetails() string {
	details := fmt.Sprintf("Department: %s\nManager: %s\nEmployees:\n", d.Name, d.Manager.GetDetails())
	for _, employee := range d.Employees {
		details += employee.GetDetails() + "\n"
	}
	return details
}
