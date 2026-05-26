# Terraform Go Demo

A comprehensive Go project demonstrating the main features of Terraform through practical examples and programmatic interaction.

## Overview

This project showcases:
- **Terraform Execution**: Run Terraform commands programmatically using Go
- **State Parsing**: Parse and analyze Terraform state files
- **Plan Analysis**: Parse and interpret Terraform execution plans
- **Custom Providers**: Example skeleton for writing custom Terraform providers
- **Terraform Concepts**: Real Terraform configurations demonstrating core features

## Features Demonstrated

### Terraform Core Concepts

1. **Resources**
   - Managed resources (create, read, update, delete)
   - Resource dependencies
   - Resource meta-arguments (count, for_each, depends_on)
   
2. **Variables**
   - Input variables with different types (string, number, bool, list, map, object, set, tuple)
   - Variable validation
   - Default values
   - Sensitive variables
   
3. **Outputs**
   - Simple outputs
   - Sensitive outputs
   - Complex outputs (objects, lists, maps)
   - Output preconditions
   
4. **Data Sources**
   - Reading existing resources
   - Using data sources in configuration
   
5. **Locals**
   - Local values and computed expressions
   - Reusing common values
   
6. **Modules**
   - Creating reusable modules
   - Module inputs and outputs
   - Calling modules with different configurations
   
7. **Lifecycle Rules**
   - `create_before_destroy`
   - `prevent_destroy`
   - `ignore_changes`
   - `replace_triggered_by`
   - Preconditions and postconditions
   
8. **Provisioners**
   - `local-exec` provisioner
   - `remote-exec` provisioner (example)
   - `file` provisioner (example)
   - Destroy-time provisioners
   
9. **Backend Configuration**
   - Local backend
   - Remote backends (S3, Azure, GCS, Terraform Cloud)
   - State locking

## Project Structure

```
terraform/
├── main.go                      # Main entry point
├── terraform_exec.go            # Terraform command execution
├── state_parser.go              # State file parsing
├── plan_parser.go               # Plan file parsing
├── provider_example.go          # Custom provider skeleton
├── go.mod                       # Go module definition
├── README.md                    # This file
│
├── examples/
│   ├── basic-infrastructure/    # Basic Terraform configuration
│   │   ├── main.tf             # Main configuration
│   │   ├── variables.tf        # Variable definitions
│   │   ├── outputs.tf          # Output definitions
│   │   └── terraform.tfvars    # Variable values
│   │
│   ├── modules/                 # Module examples
│   │   ├── web-server/         # Reusable web server module
│   │   │   └── main.tf
│   │   └── using-modules.tf    # Using modules
│   │
│   └── advanced-features/       # Advanced Terraform features
│       ├── backend.tf          # Backend configurations
│       ├── provisioners.tf     # Provisioner examples
│       └── lifecycle.tf        # Lifecycle rules
```

## Prerequisites

- Go 1.21 or later
- Terraform 1.0 or later (for running Terraform commands)

## Installation

1. Clone the repository:
```bash
cd /Users/jetbrains/myProjects/demos/terraform
```

2. Install dependencies:
```bash
go mod download
```

3. Verify Terraform installation (optional):
```bash
terraform version
```

## Usage

The project provides several commands to demonstrate different aspects of Terraform:

### 1. Execute Terraform Commands

Run Terraform commands programmatically:

```bash
go run . exec
```

This demonstrates:
- Initializing Terraform (`terraform init`)
- Validating configuration (`terraform validate`)
- Creating execution plans (`terraform plan`)
- Getting version information (`terraform version`)
- Format checking (`terraform fmt`)

### 2. Parse Terraform State

Parse and analyze a Terraform state file:

```bash
go run . state
```

This demonstrates:
- Reading state files
- Extracting resource information
- Displaying outputs
- Analyzing state structure

### 3. Parse Terraform Plans

Parse and analyze execution plans:

```bash
go run . plan
```

This demonstrates:
- Creating plan files
- Parsing plan JSON
- Analyzing resource changes
- Counting actions (create, update, delete)
- Examining output changes

### 4. Custom Provider Example

View the structure of a custom Terraform provider:

```bash
go run . provider
```

This demonstrates:
- Provider schema definition
- Resource schemas (CRUD operations)
- Data source schemas
- Provider configuration

## Terraform Examples

### Basic Infrastructure

Navigate to the basic infrastructure example:

```bash
cd examples/basic-infrastructure
```

Initialize and apply:

```bash
terraform init
terraform plan
terraform apply
```

This creates:
- Configuration files with JSON data
- Multiple files using `count`
- Environment-specific files using `for_each`
- Demonstrates variables, outputs, and data sources

### Using Modules

Navigate to the modules example:

```bash
cd examples/modules
```

Run Terraform:

```bash
terraform init
terraform plan
terraform apply
```

This demonstrates:
- Creating reusable modules
- Calling modules multiple times with different configurations
- Module outputs

### Advanced Features

Explore advanced Terraform features:

```bash
cd examples/advanced-features
```

Each file demonstrates specific features:
- `backend.tf` - Different backend configurations
- `lifecycle.tf` - Lifecycle meta-arguments
- `provisioners.tf` - Provisioner examples

## Key Terraform Features Explained

### 1. Resources

Resources are the most important element in Terraform. They represent infrastructure objects.

```hcl
resource "local_file" "example" {
  filename = "./example.txt"
  content  = "Hello, Terraform!"
}
```

### 2. Variables

Variables allow you to parameterize your configuration:

```hcl
variable "environment" {
  description = "Environment name"
  type        = string
  default     = "dev"
  
  validation {
    condition     = contains(["dev", "staging", "prod"], var.environment)
    error_message = "Must be dev, staging, or prod"
  }
}
```

### 3. Outputs

Outputs expose values from your configuration:

```hcl
output "instance_id" {
  description = "The ID of the instance"
  value       = resource.example.id
}
```

### 4. Data Sources

Data sources read information from existing resources:

```hcl
data "local_file" "config" {
  filename = "./config.json"
}
```

### 5. Locals

Locals define computed values:

```hcl
locals {
  common_tags = {
    Environment = var.environment
    ManagedBy   = "Terraform"
  }
}
```

### 6. Count and For Each

Meta-arguments for creating multiple similar resources:

```hcl
# Using count
resource "local_file" "files" {
  count    = 3
  filename = "./file-${count.index}.txt"
  content  = "File number ${count.index}"
}

# Using for_each
resource "local_file" "env_files" {
  for_each = toset(["dev", "staging", "prod"])
  filename = "./config-${each.key}.txt"
  content  = "Environment: ${each.key}"
}
```

### 7. Lifecycle Rules

Control resource behavior:

```hcl
resource "example" "resource" {
  # ...
  
  lifecycle {
    create_before_destroy = true
    prevent_destroy       = false
    ignore_changes        = [tags]
  }
}
```

### 8. Modules

Reusable infrastructure components:

```hcl
module "web_server" {
  source = "./modules/web-server"
  
  server_name = "my-server"
  port        = 8080
}
```

## Go Libraries Used

- **`github.com/hashicorp/terraform-exec`** - Execute Terraform CLI commands from Go
- **`github.com/hashicorp/terraform-json`** - Parse Terraform JSON output (state, plan)
- **`github.com/hashicorp/terraform-plugin-sdk/v2`** - Build custom Terraform providers

## Learning Resources

- [Terraform Documentation](https://www.terraform.io/docs)
- [Terraform Registry](https://registry.terraform.io/)
- [Writing Custom Providers](https://www.terraform.io/docs/extend/writing-custom-providers.html)
- [Terraform Plugin SDK](https://github.com/hashicorp/terraform-plugin-sdk)

## Common Terraform Commands

```bash
# Initialize working directory
terraform init

# Validate configuration
terraform validate

# Format configuration files
terraform fmt

# Create execution plan
terraform plan

# Apply changes
terraform apply

# Show current state
terraform show

# List resources in state
terraform state list

# Display outputs
terraform output

# Destroy infrastructure
terraform destroy
```

## Terraform Workflow

1. **Write** - Author infrastructure as code
2. **Init** - Initialize working directory
3. **Plan** - Preview changes
4. **Apply** - Create infrastructure
5. **Modify** - Update configuration
6. **Plan** - Preview updates
7. **Apply** - Update infrastructure
8. **Destroy** - Remove infrastructure (when needed)

## Best Practices Demonstrated

1. **Variable Validation** - Ensure inputs meet requirements
2. **Output Descriptions** - Document what outputs represent
3. **Module Reusability** - DRY principle for infrastructure
4. **State Management** - Proper backend configuration
5. **Resource Dependencies** - Explicit and implicit dependencies
6. **Lifecycle Management** - Control resource behavior
7. **Sensitive Data** - Mark sensitive variables and outputs
8. **Documentation** - Comment and describe resources

## Troubleshooting

### Terraform Not Found

If you see "terraform binary not found", make sure Terraform is installed:

```bash
# macOS (Homebrew)
brew install terraform

# Linux
wget https://releases.hashicorp.com/terraform/1.6.0/terraform_1.6.0_linux_amd64.zip
unzip terraform_1.6.0_linux_amd64.zip
sudo mv terraform /usr/local/bin/
```

### State File Issues

If you encounter state file issues:

```bash
# Remove state
rm terraform.tfstate

# Reinitialize
terraform init
```

### Module Not Found

If modules aren't found:

```bash
# Initialize to download modules
terraform init
```

## Contributing

This is a demonstration project. Feel free to explore and modify the examples to learn more about Terraform!

## License

This project is provided as-is for educational purposes.

