package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-exec/tfexec"
	"log"
)

func main() {
	// Set the Terraform binary path
	terraformPath := "/opt/homebrew/Cellar/terraform/1.8.5/bin/terraform" // Ensure this path points to your Terraform binary

	// Set the working directory for Terraform
	workingDir := "./terraform"

	// Initialize the Terraform executor
	tf, err := tfexec.NewTerraform(workingDir, terraformPath)
	if err != nil {
		log.Fatalf("error initializing Terraform: %v", err)
	}

	// Initialize Terraform
	if err := tf.Init(context.Background(), tfexec.Upgrade(true)); err != nil {
		log.Fatalf("error running Terraform init: %v", err)
	}

	// Apply the Terraform configuration
	if err := tf.Apply(context.Background()); err != nil {
		log.Fatalf("error running Terraform apply: %v", err)
	}

	fmt.Println("Terraform apply completed successfully!")
}
