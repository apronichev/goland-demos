package main

import "fmt"

func case1() {
	user := NewUser("", 21)
	fmt.Printf("Age: %d, %s", user.Age, user.Name)
}

func case2() {
	user, err := NewUserMoreCorrect("", -10)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	fmt.Printf("Age: %d, %s", user.Age, user.Name)
}

func main() {
	case1()
}
