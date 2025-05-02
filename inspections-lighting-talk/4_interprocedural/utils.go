package main

import "errors"

type User struct {
	Name string
	Age  int
}

func isValidName(name string) bool {
	return name != ""
}

func isValidAge(age int) bool {
	return age >= 0
}

func NewUser(name string, age int) *User {
	if !isValidName(name) {
		return nil
	}
	return &User{Name: name, Age: age}
}

func NewUserMoreCorrect(name string, age int) (*User, error) {
	if !isValidName(name) {
		return nil, errors.New("invalid name")
	}
	if !isValidAge(age) {
		return nil, nil
	}
	return &User{Name: name}, nil
}
