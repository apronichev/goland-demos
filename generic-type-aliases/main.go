// File: myapp/main_test.go
package main

import (
	"fmt"
	"generic-type-aliases/wrapper"
)

type Player struct {
	Name string
	HP   int
}

func main() {
	// Using the re-exported type
	playerLoc := wrapper.Location[Player]{
		Object: Player{Name: "Hero", HP: 100},
		X:      10.5,
		Y:      20.3,
	}

	fmt.Printf("Player at (%.1f, %.1f): %s\n", playerLoc.X, playerLoc.Y, playerLoc.Object.Name)
}
