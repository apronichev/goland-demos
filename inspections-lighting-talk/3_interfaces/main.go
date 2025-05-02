package main

import (
	"fmt"
)

type InvalidUsernameError struct {
	Username string
	Reason   string
}

func (e *InvalidUsernameError) Error() string {
	return fmt.Sprintf("Username %s is invalid: %s", e.Username, e.Reason)
}

func Validate(username string) error {
	var err *InvalidUsernameError = nil
	if len(username) == 0 {
		err = &InvalidUsernameError{Username: username, Reason: "Empty username is not allowed"}
	}
	if username == "😞" {
		err = &InvalidUsernameError{Username: username, Reason: "Sadness is forbidden!"}
	}
	return err
}

func main() {
	name := ""
	err := Validate(name)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	fmt.Print("Username is valid!")
}
