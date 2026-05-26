# Terraform Features Demonstrated

This project provides comprehensive examples of Terraform's main features through both Go code and Terraform configurations.

## Core Terraform Concepts

### 1. **Resources** 🔧
Resources are the building blocks of Terraform configurations, representing infrastructure objects.

**Files:** `examples/basic-infrastructure/main.tf`

**Examples:**
- Creating resources with `resource` blocks
- Using `count` to create multiple instances
- Using `for_each` for map-based resource creation
- Resource dependencies with `depends_on`

```hcl
resource "local_file" "example" {
  filename = "./example.txt"
  content  = "Hello, Terraform!"
}
```

### 2. **Variables** 📝
Input variables allow parameterization of configurations.

**Files:** `examples/basic-infrastructure/variables.tf`

**Types Demonstrated:**
- `string` - Text values
- `number` - Numeric values
- `bool` - Boolean true/false
- `list` - Ordered collections
- `map` - Key-value pairs
- `object` - Structured data
- `tuple` - Fixed-length lists with types
- `set` - Unordered unique collections

**Features:**
- Default values
- Variable validation rules
- Sensitive variables (hidden in output)
- Nullable variables
- Type constraints

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

### 3. **Outputs** 📤
Export values from your configuration for use by other modules or display to users.

**Files:** `examples/basic-infrastructure/outputs.tf`

**Features:**
- Simple value outputs
- Sensitive outputs (masked in logs)
- Complex outputs (objects, lists, maps)
- Output preconditions
- Outputs with `depends_on`

```hcl
output "instance_id" {
  description = "The ID of the instance"
  value       = resource.example.id
  sensitive   = false
}
```

### 4. **Data Sources** 🔍
Read information from existing infrastructure or external sources.

**Files:** `examples/basic-infrastructure/main.tf`

**Examples:**
- Reading existing files
- Querying cloud resources
- Using data source outputs in resources

```hcl
data "local_file" "config" {
  filename = "./config.json"
}
```

### 5. **Locals** 💡
Define computed values and reusable expressions.

**Files:** `examples/basic-infrastructure/main.tf`

**Use Cases:**
- Computing values from variables
- Creating reusable expressions
- Building complex data structures
- Timestamp and dynamic values

```hcl
locals {
  common_tags = {
    Environment = var.environment
    ManagedBy   = "Terraform"
  }
}
```

### 6. **Meta-Arguments** 🔄

#### Count
Create multiple similar resources with an index.

```hcl
resource "local_file" "files" {
  count    = 3
  filename = "./file-${count.index}.txt"
  content  = "File ${count.index}"
}
```

#### For Each
Create resources from a map or set.

```hcl
resource "local_file" "env_files" {
  for_each = toset(["dev", "staging", "prod"])
  filename = "./config-${each.key}.txt"
  content  = "Environment: ${each.key}"
}
```

#### Depends On
Explicitly declare dependencies.

```hcl
resource "example" "dependent" {
  # ...
  depends_on = [other_resource.example]
}
```

### 7. **Modules** 📦
Reusable infrastructure components.

**Files:** `examples/modules/`

**Features:**
- Module inputs (variables)
- Module outputs
- Module composition
- Calling modules multiple times
- Local and remote modules

**Example:**
```hcl
module "web_server" {
  source      = "./modules/web-server"
  server_name = "my-server"
  port        = 8080
}
```

### 8. **Lifecycle Rules** ♻️

**Files:** `examples/advanced-features/lifecycle.tf`

**Rules:**
- `create_before_destroy` - Create replacement before destroying
- `prevent_destroy` - Prevent accidental destruction
- `ignore_changes` - Ignore changes to specific attributes
- `replace_triggered_by` - Force replacement on change
- `precondition` - Validate before operations
- `postcondition` - Validate after operations

```hcl
resource "example" "resource" {
  # ...
  
  lifecycle {
    create_before_destroy = true
    ignore_changes        = [tags]
    
    precondition {
      condition     = var.size > 0
      error_message = "Size must be positive"
    }
  }
}
```

### 9. **Provisioners** 🛠️

**Files:** `examples/advanced-features/provisioners.tf`

**Types:**
- `local-exec` - Run commands locally
- `remote-exec` - Run commands on remote resources
- `file` - Copy files to remote resources

**Features:**
- Destroy-time provisioners
- Failure handling (`continue` or `fail`)
- Environment variables
- Working directory specification

```hcl
resource "example" "resource" {
  # ...
  
  provisioner "local-exec" {
    command = "echo 'Resource created'"
    
    environment = {
      VAR_NAME = "value"
    }
  }
}
```

### 10. **Backend Configuration** 💾

**Files:** `examples/advanced-features/backend.tf`

**Backends Demonstrated:**
- Local (default) - Store state in local file
- S3 - AWS S3 with DynamoDB locking
- Azure Storage - Azure blob storage
- GCS - Google Cloud Storage
- Terraform Cloud - HashiCorp managed state

```hcl
terraform {
  backend "s3" {
    bucket         = "my-terraform-state"
    key            = "terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-locks"
  }
}
```

### 11. **Expressions and Functions** 🧮

**Demonstrated in:** Multiple files

**String Functions:**
- `format()`, `join()`, `split()`, `replace()`

**Collection Functions:**
- `length()`, `concat()`, `merge()`, `flatten()`
- `keys()`, `values()`, `lookup()`
- `element()`, `contains()`, `distinct()`

**Type Conversion:**
- `tostring()`, `tonumber()`, `tobool()`
- `tolist()`, `tomap()`, `toset()`

**Encoding Functions:**
- `jsonencode()`, `jsondecode()`
- `base64encode()`, `base64decode()`

**Special Functions:**
- `timestamp()` - Current timestamp
- `uuid()` - Generate UUID
- `file()` - Read file contents
- `path.module`, `path.root`, `path.cwd` - Path references

### 12. **State Management** 📊

**Go Files:** `state_parser.go`

**Features:**
- State file structure
- Resource tracking
- Output values storage
- State versioning
- State locking (with backends)

### 13. **Planning** 📋

**Go Files:** `plan_parser.go`

**Features:**
- Execution plan creation
- Change preview (create, update, delete)
- Resource dependencies
- Output changes
- Plan file format

## Go Integration Features

### 1. **Terraform Execution** (`terraform_exec.go`)
- Initialize Terraform
- Validate configurations
- Create plans
- Apply changes
- Get version info
- Format checking

### 2. **State Parsing** (`state_parser.go`)
- Read state files
- Parse JSON structure
- Extract resource information
- Analyze state content

### 3. **Plan Parsing** (`plan_parser.go`)
- Create plan files
- Parse plan JSON
- Analyze resource changes
- Count actions by type

### 4. **Custom Providers** (`provider_example.go`)
- Provider schema definition
- Resource CRUD operations
- Data source read operations
- Provider configuration

## Terraform Best Practices Demonstrated

✅ **Use Variables** - Parameterize configurations
✅ **Validate Inputs** - Add validation rules to variables
✅ **Document Outputs** - Add descriptions to all outputs
✅ **Use Modules** - Create reusable components
✅ **State Management** - Use remote backends for teams
✅ **Version Pinning** - Pin provider versions
✅ **Resource Dependencies** - Use depends_on when needed
✅ **Lifecycle Rules** - Prevent accidental destruction
✅ **Sensitive Data** - Mark sensitive values appropriately
✅ **Code Formatting** - Use `terraform fmt`
✅ **Validation** - Run `terraform validate` regularly

## Quick Reference

### Variable Types
| Type | Example | Description |
|------|---------|-------------|
| `string` | `"hello"` | Text values |
| `number` | `42` | Numeric values |
| `bool` | `true` | Boolean values |
| `list(type)` | `["a", "b"]` | Ordered collection |
| `map(type)` | `{a = "x"}` | Key-value pairs |
| `set(type)` | `["a", "b"]` | Unique values |
| `object({...})` | Complex | Structured data |
| `tuple([...])` | `["a", 1]` | Fixed-length list |

### Meta-Arguments
| Argument | Purpose |
|----------|---------|
| `count` | Create multiple instances with index |
| `for_each` | Create instances from set/map |
| `depends_on` | Explicit dependencies |
| `provider` | Select provider configuration |
| `lifecycle` | Control resource behavior |

### Lifecycle Rules
| Rule | Purpose |
|------|---------|
| `create_before_destroy` | Replace resources safely |
| `prevent_destroy` | Prevent accidental deletion |
| `ignore_changes` | Ignore drift in attributes |
| `replace_triggered_by` | Force replacement on change |

### Common Functions
| Function | Purpose |
|----------|---------|
| `file(path)` | Read file contents |
| `jsonencode(value)` | Convert to JSON |
| `timestamp()` | Current timestamp |
| `length(collection)` | Count items |
| `concat(lists...)` | Combine lists |
| `merge(maps...)` | Combine maps |
| `lookup(map, key, default)` | Get map value |
| `contains(list, value)` | Check membership |

## Project Structure Summary

```
terraform/
├── Go Source Files (Terraform Integration)
│   ├── main.go                  - Entry point
│   ├── terraform_exec.go        - Execute Terraform
│   ├── state_parser.go          - Parse state files
│   ├── plan_parser.go           - Parse plan files
│   └── provider_example.go      - Provider skeleton
│
├── Terraform Examples (Real Configurations)
│   ├── basic-infrastructure/    - Core features
│   │   ├── main.tf             - Resources, data, locals
│   │   ├── variables.tf        - All variable types
│   │   ├── outputs.tf          - Output examples
│   │   └── terraform.tfvars    - Variable values
│   │
│   ├── modules/                 - Reusable modules
│   │   ├── web-server/         - Example module
│   │   └── using-modules.tf    - Module usage
│   │
│   └── advanced-features/       - Advanced concepts
│       ├── backend.tf          - State backends
│       ├── lifecycle.tf        - Lifecycle rules
│       └── provisioners.tf     - Provisioner examples
│
└── Documentation
    ├── README.md               - Full documentation
    ├── QUICKSTART.md          - Getting started
    └── FEATURES.md            - This file
```

## Learning Path

1. **Start Here:** `examples/basic-infrastructure/main.tf`
   - Understand resources, variables, outputs
   
2. **Next:** `examples/basic-infrastructure/variables.tf`
   - Learn all variable types and validation
   
3. **Then:** `examples/modules/`
   - Create and use reusable modules
   
4. **Advanced:** `examples/advanced-features/`
   - Lifecycle rules, provisioners, backends
   
5. **Go Integration:** Run Go examples
   - Execute: `./terraform-demo exec`
   - State: `./terraform-demo state`
   - Plan: `./terraform-demo plan`
   - Provider: `./terraform-demo provider`

## Resources

- 📖 [Terraform Documentation](https://www.terraform.io/docs)
- 🎓 [HashiCorp Learn](https://learn.hashicorp.com/terraform)
- 📦 [Terraform Registry](https://registry.terraform.io/)
- 🔧 [Provider Development](https://www.terraform.io/docs/extend)
- 💬 [Terraform Community Forum](https://discuss.hashicorp.com/c/terraform-core)

