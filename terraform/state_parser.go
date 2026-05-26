package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	tfjson "github.com/hashicorp/terraform-json"
)

// runStateParser demonstrates parsing and working with Terraform state files
func runStateParser() {
	fmt.Println("=== Terraform State Parser Example ===")
	fmt.Println("Demonstrating how to parse and analyze Terraform state files")
	fmt.Println()

	statePath := "./examples/basic-infrastructure/terraform.tfstate"

	// Check if state file exists
	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		fmt.Println("State file not found. Creating example state structure...")
		printExampleStateStructure()
		return
	}

	// Read state file
	data, err := os.ReadFile(statePath)
	if err != nil {
		log.Fatalf("Error reading state file: %v", err)
	}

	// Parse state
	var state tfjson.State
	err = json.Unmarshal(data, &state)
	if err != nil {
		log.Fatalf("Error parsing state: %v", err)
	}

	// Display state information
	fmt.Printf("Terraform Version: %s\n", state.TerraformVersion)
	fmt.Printf("Format Version: %d\n", state.FormatVersion)
	fmt.Println()

	// Display resources
	if state.Values != nil && state.Values.RootModule != nil {
		fmt.Println("Resources in State:")
		fmt.Println("-------------------")
		displayModuleResources(state.Values.RootModule, "")
	}

	// Display outputs
	if len(state.Values.Outputs) > 0 {
		fmt.Println("\nOutputs:")
		fmt.Println("--------")
		for name, output := range state.Values.Outputs {
			fmt.Printf("  %s = %v (sensitive: %t)\n", name, output.Value, output.Sensitive)
		}
	}

	fmt.Println("\n=== State Parser Complete ===")
}

// displayModuleResources recursively displays resources in a module
func displayModuleResources(module *tfjson.StateModule, indent string) {
	// Display resources in current module
	for _, resource := range module.Resources {
		fmt.Printf("%s- %s (%s)\n", indent, resource.Address, resource.Type)
		fmt.Printf("%s  Provider: %s\n", indent, resource.ProviderName)
		if len(resource.AttributeValues) > 0 {
			fmt.Printf("%s  Attributes: %d\n", indent, len(resource.AttributeValues))
		}
	}

	// Display child modules
	for _, childModule := range module.ChildModules {
		fmt.Printf("%s\nModule: %s\n", indent, childModule.Address)
		displayModuleResources(childModule, indent+"  ")
	}
}

// printExampleStateStructure shows what a Terraform state looks like
func printExampleStateStructure() {
	fmt.Println("Terraform State Structure:")
	fmt.Println("--------------------------")
	fmt.Println(`
A Terraform state file contains:

1. Version Information:
   - Terraform version used
   - State format version

2. Resources:
   - Resource type (e.g., aws_instance, google_compute_instance)
   - Resource name
   - Provider configuration
   - Current attribute values
   - Dependencies

3. Outputs:
   - Output values defined in configuration
   - Sensitivity flags

4. Metadata:
   - Serial number (increments with each change)
   - Lineage (unique identifier for the state)

Example resource in state:
{
  "address": "local_file.example",
  "mode": "managed",
  "type": "local_file",
  "name": "example",
  "provider_name": "local",
  "attributes": {
    "filename": "./output.txt",
    "content": "Hello, Terraform!"
  }
}

To generate a state file, run: terraform apply in ./examples/basic-infrastructure/
	`)
}
