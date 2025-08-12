package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// Create a Root for secure file operations within 'tmp/'
	root, err := os.OpenRoot("tmp")
	if err != nil {
		log.Fatal(err)
	}
	defer root.Close()

	// Create directory structure
	fmt.Println("Creating directories...")
	if err := root.MkdirAll("app/config", 0755); err != nil {
		log.Fatal(err)
	}

	// Write configuration file
	fmt.Println("Writing config file...")
	config := `{
  "app": "myapp",
  "version": "1.0"
}`
	if err := root.WriteFile("app/config/settings.json", []byte(config), 0644); err != nil {
		log.Fatal(err)
	}

	// Read the file back
	fmt.Println("Reading config file...")
	data, err := root.ReadFile("app/config/settings.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Config content: %s\n", data)

	// Create a symlink
	fmt.Println("Creating symlink...")
	if err := root.Symlink("app/config/settings.json", "current-config.json"); err != nil {
		log.Fatal(err)
	}

	// Read the symlink target
	target, err := root.Readlink("current-config.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Symlink points to: %s\n", target)

	// Change file permissions
	fmt.Println("Changing permissions...")
	if err := root.Chmod("app/config/settings.json", 0600); err != nil {
		log.Fatal(err)
	}

	// Update file timestamps
	fmt.Println("Updating timestamps...")
	now := time.Now()
	if err := root.Chtimes("app/config/settings.json", now, now); err != nil {
		log.Fatal(err)
	}

	// Create a hard link
	fmt.Println("Creating hard link...")
	if err := root.Link("app/config/settings.json", "backup-config.json"); err != nil {
		log.Fatal(err)
	}

	// Rename a file
	fmt.Println("Renaming file...")
	if err := root.Rename("backup-config.json", "settings-backup.json"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("File operations completed successfully!")
}
