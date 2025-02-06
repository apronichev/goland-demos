package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

func main() {
	j := `
    {
      "email": "noah.poulsen@example.com",
      "gender": "male",
      "first_name": "Robert",
      "last_name": "Paulson",
      "location": {
        "street": "7998 oddenvej",
        "city": "nr åby",
        "state": "syddanmark",
        "postcode": 73617
      },
      "username": "purplesnake503",
      "password": "myZelda@",
      "picture": "img/41.jpg"
    }
`

	registrationDate, _ := time.Parse("2006-01", "2022-11")
	tellUserPassword(j, registrationDate)
}

func tellUserPassword(userData string, t time.Time) {
	var user T
	err := json.Unmarshal([]byte(userData), &user)
	if err != nil {
		return
	}
	fmt.Println(verifyPassword(user.Password))
	fmt.Println("%s %s's password: %s. The '%s' email is %s. Account created: %s", user.FirstName, user.LastName, user.Password, user.Email, IsValidEmail(user.Email), t.Local())
}

func verifyPassword(s string) (sixOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range s {
		if IsNumber(c) {
			number = true
		} else if IsUpper(c) {
			upper = true
			letters++
		} else if IsSymbol(c) || IsPunct(c) {
			special = true
		} else if IsLetter(c) {
			letters++
		} else {
			return false, false, false, false
		}
	}
	sixOrMore = letters >= 6
	return
}

func IsValidEmail(email string) string {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(email) {
		return "valid"
	}
	return "invalid"
}
