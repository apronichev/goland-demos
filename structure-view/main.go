package main

import (
	"fmt"
	"structure-view/company"
)

// todo using AI Assistant, refactor the code below
func main() {
	// Create employees
	employee1 := company.Employee{Name: "Alice", Age: 28, Salary: 50000, Email: "alice@example.com"}
	employee2 := company.Employee{Name: "Bob", Age: 34, Salary: 60000, Email: "bob@example.com"}

	// Validate employee emails
	fmt.Printf("Is %s's email valid? %v\n", employee1.Name, employee1.IsValidEmail())
	fmt.Printf("Is %s's email valid? %v\n", employee2.Name, employee2.IsValidEmail())

	// Create a manager
	manager := &company.Manager{
		Employee:   company.Employee{Name: "Charlie", Age: 40, Salary: 80000, Email: "charlie@example.com"},
		Department: "Engineering",
	}

	// Validate manager email
	fmt.Printf("Is %s's email valid? %v\n", manager.Name, manager.IsValidEmail())

	// Create a department and add employees
	department := company.Department{Name: "Engineering", Manager: manager}
	department.AddEmployee(employee1)
	department.AddEmployee(employee2)

	// Print details
	fmt.Println(department.GetDetails())
}
