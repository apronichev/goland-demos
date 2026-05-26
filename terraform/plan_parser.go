package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
)

// runPlanParser demonstrates parsing and analyzing Terraform plans
func runPlanParser() {
	fmt.Println("=== Terraform Plan Parser Example ===")
	fmt.Println("Demonstrating how to parse and analyze Terraform execution plans")
	fmt.Println()

	workDir := "./examples/basic-infrastructure"

	// Find terraform binary
	terraformPath, err := findTerraformBinary()
	if err != nil {
		fmt.Println("Terraform binary not found. Showing example plan structure...")
		printExamplePlanStructure()
		return
	}

	// Create terraform instance
	tf, err := tfexec.NewTerraform(workDir, terraformPath)
	if err != nil {
		log.Fatalf("Error creating terraform instance: %v", err)
	}

	ctx := context.Background()

	// Initialize if needed
	if _, err := os.Stat(workDir + "/.terraform"); os.IsNotExist(err) {
		fmt.Println("Initializing Terraform...")
		if err := tf.Init(ctx); err != nil {
			log.Fatalf("Error initializing: %v", err)
		}
	}

	// Create a plan file
	planPath := workDir + "/tfplan"
	fmt.Println("Creating plan...")
	_, err = tf.Plan(ctx, tfexec.Out(planPath))
	if err != nil {
		log.Fatalf("Error creating plan: %v", err)
	}

	// Show the plan in JSON format
	fmt.Println("Parsing plan...")
	plan, err := tf.ShowPlanFile(ctx, planPath)
	if err != nil {
		log.Fatalf("Error parsing plan: %v", err)
	}

	// Analyze the plan
	analyzePlan(plan)

	// Clean up plan file
	os.Remove(planPath)

	fmt.Println("\n=== Plan Parser Complete ===")
}

// analyzePlan analyzes and displays information about a Terraform plan
func analyzePlan(plan *tfjson.Plan) {
	fmt.Println("\nPlan Analysis:")
	fmt.Println("--------------")

	// Count changes by action
	actions := map[string]int{
		"create": 0,
		"update": 0,
		"delete": 0,
		"read":   0,
		"noop":   0,
	}

	// Analyze resource changes
	for _, rc := range plan.ResourceChanges {
		for _, action := range rc.Change.Actions {
			actionStr := string(action)
			actions[actionStr]++
		}

		// Display change details
		fmt.Printf("\n%s: %s\n", rc.Address, rc.Change.Actions)

		// Show what's changing
		if rc.Change.Before != nil && rc.Change.After != nil {
			fmt.Println("  Changes detected in attributes")
		} else if rc.Change.Before == nil {
			fmt.Println("  Will be created")
		} else if rc.Change.After == nil {
			fmt.Println("  Will be destroyed")
		}
	}

	// Summary
	fmt.Println("\nSummary:")
	fmt.Printf("  Resources to create: %d\n", actions["create"])
	fmt.Printf("  Resources to update: %d\n", actions["update"])
	fmt.Printf("  Resources to delete: %d\n", actions["delete"])
	fmt.Printf("  Resources to read:   %d\n", actions["read"])

	// Output changes
	if len(plan.OutputChanges) > 0 {
		fmt.Println("\nOutput Changes:")
		for name, change := range plan.OutputChanges {
			fmt.Printf("  %s: %v -> %v\n", name, change.Before, change.After)
		}
	}

	// Variables
	if len(plan.Variables) > 0 {
		fmt.Println("\nVariables:")
		for name, variable := range plan.Variables {
			fmt.Printf("  %s = %v\n", name, variable.Value)
		}
	}
}

// printExamplePlanStructure shows what a Terraform plan contains
func printExamplePlanStructure() {
	fmt.Println("Terraform Plan Structure:")
	fmt.Println("-------------------------")
	fmt.Println(`
A Terraform plan contains:

1. Resource Changes:
   - Actions: create, update, delete, read, no-op
   - Before values (current state)
   - After values (desired state)
   - Attribute changes

2. Output Changes:
   - New outputs
   - Modified outputs
   - Deleted outputs

3. Variables:
   - Input variable values used in the plan

4. Configuration:
   - Provider configurations
   - Resource configurations

Example Actions:
- create: New resource will be created
- update: Existing resource will be modified
- delete: Resource will be destroyed
- read: Data source will be read
- no-op: No changes needed

Example Resource Change:
{
  "address": "local_file.demo",
  "mode": "managed",
  "type": "local_file",
  "name": "demo",
  "change": {
    "actions": ["create"],
    "before": null,
    "after": {
      "filename": "./demo.txt",
      "content": "Hello from Terraform!"
    }
  }
}

To generate a plan, run: terraform plan in ./examples/basic-infrastructure/
	`)
}
