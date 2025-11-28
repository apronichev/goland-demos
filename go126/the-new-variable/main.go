package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type UserProfile struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Age      *int    `json:"age,omitempty"`
	IsActive *bool   `json:"is_active,omitempty"`
	Bio      *string `json:"bio,omitempty"`
}

func main() {
	user := createProfileNewWay()
	user2 := createProfile()

	fmt.Println("--- Generating JSON with Go 1.26 Features ---")
	jsonData, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(jsonData))

	fmt.Println("--- Generating JSON using an old approach ---")
	jsonData2, _ := json.MarshalIndent(user2, "", "  ")
	fmt.Println(string(jsonData2))
}

func createProfileNewWay() UserProfile {
	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)

	return UserProfile{
		ID:       101,
		Username: "FutureGopher",

		Age:      new(calculateAge(birthDate)),
		IsActive: new(true),
		Bio:      new("Ready for Go 1.26!"),
	}
}

func createProfile() UserProfile {
	// birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	ageVal := 34
	activeVal := true

	return UserProfile{
		ID:       102,
		Username: "Gopher",
		Age:      &ageVal,
		IsActive: &activeVal,
		// Bio: new("Ready for Go 1.26!"),
	}
}

func calculateAge(born time.Time) int {
	return int(time.Since(born).Hours() / (365.25 * 24))
}
