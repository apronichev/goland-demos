package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
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
	fmt.Println("%s %s's password: %s. Account created: %s", user.FirstName, user.LastName, user.Password, t.Local())
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
