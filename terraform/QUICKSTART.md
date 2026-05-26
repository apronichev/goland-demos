# Quick Start Guide

## Running the Go Examples

### 1. Build the project

```bash
go build -o terraform-demo
```

### 2. Run different examples

#### Execute Terraform Commands
```bash
./terraform-demo exec
```

This will demonstrate:
- Terraform initialization
- Configuration validation
- Plan creation
- Version checking
- Format checking

#### Parse Terraform State
```bash
./terraform-demo state
```

This will show the structure of Terraform state files and explain what they contain.

#### Parse Terraform Plans
```bash
./terraform-demo plan
```

This will create and parse a Terraform plan file, showing all changes.

#### View Custom Provider Example
```bash
./terraform-demo provider
```

This shows the structure of a custom Terraform provider written in Go.

## Running the Terraform Examples

### Basic Infrastructure Example

```bash
cd examples/basic-infrastructure
terraform init
terraform plan
terraform apply
```

This creates:
- Random strings and UUIDs
- Local configuration files in JSON format
- Multiple files using `count`
- Environment-specific files using `for_each`

To see the outputs:
```bash
terraform output
```

To destroy:
```bash
terraform destroy
```

### Module Example

```bash
cd examples/modules
terraform init
terraform plan
terraform apply
```

This demonstrates:
- Creating reusable modules
- Calling the same module multiple times with different configurations
- Module outputs

### Advanced Features

Individual Terraform files in `examples/advanced-features/` demonstrate:
- **backend.tf** - Different backend configurations (S3, Azure, GCS, etc.)
- **lifecycle.tf** - Lifecycle rules (create_before_destroy, ignore_changes, etc.)
- **provisioners.tf** - Local-exec, remote-exec, and file provisioners

To try these:
```bash
cd examples/advanced-features
terraform init
terraform plan -target=local_file.cbd_file
```

## Key Terraform Commands

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

## Terraform Features Demonstrated

✅ **Resources** - Create, read, update, delete infrastructure
✅ **Variables** - Parameterize configurations with different types
✅ **Outputs** - Export values from configurations
✅ **Data Sources** - Read existing resources
✅ **Locals** - Define computed values
✅ **Count** - Create multiple similar resources
✅ **For Each** - Create resources from a set
✅ **Modules** - Reusable infrastructure components
✅ **Lifecycle Rules** - Control resource behavior
✅ **Provisioners** - Execute scripts during resource lifecycle
✅ **Backend Configuration** - Remote state storage
✅ **State Management** - Track infrastructure state
✅ **Plan Analysis** - Preview changes before applying
✅ **Variable Validation** - Ensure inputs meet requirements
✅ **Sensitive Values** - Handle secrets securely

## Go + Terraform Integration

This project demonstrates:
- Using `terraform-exec` to run Terraform from Go
- Parsing Terraform state files with `terraform-json`
- Analyzing execution plans programmatically
- Creating custom Terraform providers with the Plugin SDK

## Next Steps

1. Explore the example configurations in `examples/`
2. Modify variables in `terraform.tfvars`
3. Try creating your own resources
4. Experiment with modules
5. Study the Go code to understand Terraform's API

## Troubleshooting

**Terraform not found?**
```bash
# macOS
brew install terraform

# Or download from: https://www.terraform.io/downloads
```

**State locked?**
```bash
# Force unlock (use with caution)
terraform force-unlock <lock-id>
```

**Need to start fresh?**
```bash
rm -rf .terraform terraform.tfstate*
terraform init
```

## Resources

- [Terraform Documentation](https://www.terraform.io/docs)
- [Terraform Registry](https://registry.terraform.io/)
- [HashiCorp Learn](https://learn.hashicorp.com/terraform)

