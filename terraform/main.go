package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "exec":
		// Execute Terraform commands programmatically
		runTerraformExec()
	case "state":
		// Parse and display Terraform state
		runStateParser()
	case "plan":
		// Parse and display Terraform plan
		runPlanParser()
	case "provider":
		// Run custom provider example
		runProviderExample()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Terraform Go Demo - Demonstrates Terraform features using Go")
	fmt.Println("\nUsage:")
	fmt.Println("  go run main.go <command>")
	fmt.Println("\nCommands:")
	fmt.Println("  exec      - Execute Terraform commands programmatically")
	fmt.Println("  state     - Parse and display Terraform state")
	fmt.Println("  plan      - Parse and display Terraform plan")
	fmt.Println("  provider  - Run custom provider example")
	fmt.Println("\nExamples:")
	fmt.Println("  go run main.go exec")
	fmt.Println("  go run main.go state")
}
