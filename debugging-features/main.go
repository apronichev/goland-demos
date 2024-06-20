package main

import (
	"encoding/json"
	"fmt"
	"time"
	"unicode"
)

// todo copy the JSON fragment below and paste it here, name struct as T

func main() {
	//todo inject JSON
	userJSON := `
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
      "password": "Zelda@",
      "picture": "img/41.jpg"
    }
`

	registrationDate, _ := time.Parse("2006-01", "2022-11")
	tellUserPassword(userJSON, registrationDate)
}

func tellUserPassword(userData string, t time.Time) {
	var user T
	err := json.Unmarshal([]byte(userData), &user)
	if err != nil {
		return
	}
	fmt.Println(verifyPassword(user.Password))
	//todo quick-fix to Printf
	fmt.Println("%s %s's password: %s. Account created: %s", user.FirstName, user.LastName, user.Password, t.Local())
}

// todo while debugging press Alt+F8 for Evaluate expression. Show watches
func verifyPassword(s string) (sixOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range s {
		if isNumber(c) {
			number = true
		} else if isUpper(c) {
			upper = true
			letters++
			//todo put a breakpoint on the line below, run 'Debug', make 'Smart step into'
		} else if isSymbol(c) || isPunct(c) {
			special = true
		} else if isLetter(c) {
			letters++
		} else {
			return false, false, false, false
		}
	}
	sixOrMore = letters >= 6
	return
}

func isNumber(c rune) bool {
	return unicode.IsNumber(c)
}

func isUpper(c rune) bool {
	return unicode.IsUpper(c)
}

func isPunct(c rune) bool {
	return unicode.IsPunct(c)
}

func isSymbol(c rune) bool {
	return unicode.IsSymbol(c)
}

func isLetter(c rune) bool {
	return unicode.IsLetter(c) || c == ' '
}
