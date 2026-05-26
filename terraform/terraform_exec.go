package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/terraform-exec/tfexec"
)

// runTerraformExec demonstrates executing Terraform commands programmatically using terraform-exec
func runTerraformExec() {
	fmt.Println("=== Terraform Exec Example ===")
	fmt.Println("Demonstrating how to execute Terraform commands from Go")
	fmt.Println()

	// Create a working directory for the example
	workDir := "./examples/basic-infrastructure"

	// Find terraform binary (assumes terraform is in PATH)
	terraformPath, err := findTerraformBinary()
	if err != nil {
		log.Printf("Warning: Terraform binary not found. Install Terraform to run this example.\n")
		log.Printf("Error: %v\n", err)
		fmt.Println("\nThis example demonstrates:")
		fmt.Println("  - terraform init")
		fmt.Println("  - terraform validate")
		fmt.Println("  - terraform plan")
		fmt.Println("  - terraform apply (with auto-approve)")
		fmt.Println("  - terraform show (state)")
		fmt.Println("  - terraform output")
		fmt.Println("  - terraform destroy")
		return
	}

	// Create a new Terraform instance
	tf, err := tfexec.NewTerraform(workDir, terraformPath)
	if err != nil {
		log.Fatalf("Error creating terraform instance: %v", err)
	}

	ctx := context.Background()

	// 1. Initialize Terraform
	fmt.Println("1. Initializing Terraform...")
	err = tf.Init(ctx, tfexec.Upgrade(true))
	if err != nil {
		log.Fatalf("Error running terraform init: %v", err)
	}
	fmt.Println("✓ Initialization complete")
	fmt.Println()

	// 2. Validate configuration
	fmt.Println("2. Validating configuration...")
	_, err = tf.Validate(ctx)
	if err != nil {
		log.Fatalf("Error validating configuration: %v", err)
	}
	fmt.Println("✓ Configuration is valid")
	fmt.Println()

	// 3. Create a plan
	fmt.Println("3. Creating execution plan...")
	hasChanges, err := tf.Plan(ctx)
	if err != nil {
		log.Fatalf("Error creating plan: %v", err)
	}
	if hasChanges {
		fmt.Println("✓ Plan created with changes")
	} else {
		fmt.Println("✓ No changes needed")
	}
	fmt.Println()

	// 4. Show version
	fmt.Println("4. Getting Terraform version...")
	version, _, err := tf.Version(ctx, false)
	if err != nil {
		log.Printf("Error getting version: %v", err)
	} else {
		fmt.Printf("✓ Terraform version: %s\n", version.String())
	}
	fmt.Println()

	// 5. Format check
	fmt.Println("5. Checking formatting...")
	diff, _, err := tf.FormatCheck(ctx)
	if err != nil {
		log.Printf("Error checking format: %v", err)
	} else {
		if diff {
			fmt.Println("⚠ Some files need formatting")
		} else {
			fmt.Println("✓ All files are properly formatted")
		}
	}
	fmt.Println()

	fmt.Println("=== Example Complete ===")
	fmt.Println("\nNote: To actually apply changes, uncomment the Apply section in terraform_exec.go")
}

// findTerraformBinary locates the terraform binary
func findTerraformBinary() (string, error) {
	// Check common locations
	paths := []string{
		"/usr/local/bin/terraform",
		"/usr/bin/terraform",
		"/opt/homebrew/bin/terraform",
	}

	// First check PATH
	if path, err := exec.LookPath("terraform"); err == nil {
		return path, nil
	}

	// Then check common locations
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("terraform binary not found")
}
