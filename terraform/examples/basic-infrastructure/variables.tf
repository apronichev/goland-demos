# Variables demonstrating different types and features

# String variable with validation
variable "region" {
  description = "AWS region (example, not used in this demo)"
  type        = string
  default     = "us-east-1"
  
  validation {
    condition     = can(regex("^[a-z]{2}-[a-z]+-[0-9]$", var.region))
    error_message = "Region must be a valid AWS region format (e.g., us-east-1)."
  }
}

# Number variable
variable "max_items" {
  description = "Maximum number of items"
  type        = number
  default     = 10
  
  validation {
    condition     = var.max_items > 0 && var.max_items <= 100
    error_message = "Max items must be between 1 and 100."
  }
}

# Boolean variable
variable "enable_monitoring" {
  description = "Enable monitoring features"
  type        = bool
  default     = true
}

# List variable
variable "allowed_ips" {
  description = "List of allowed IP addresses"
  type        = list(string)
  default     = ["10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"]
}

# Map variable
variable "instance_types" {
  description = "Map of environment to instance type"
  type        = map(string)
  default = {
    dev     = "t2.micro"
    staging = "t2.small"
    prod    = "t2.medium"
  }
}

# Object variable
variable "database_config" {
  description = "Database configuration object"
  type = object({
    engine         = string
    version        = string
    port           = number
    multi_az       = bool
    backup_enabled = bool
  })
  default = {
    engine         = "postgres"
    version        = "14.5"
    port           = 5432
    multi_az       = false
    backup_enabled = true
  }
}

# Tuple variable
variable "cidr_blocks" {
  description = "List of CIDR blocks with specific types"
  type        = tuple([string, string, string])
  default     = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
}

# Set variable
variable "availability_zones" {
  description = "Set of availability zones"
  type        = set(string)
  default     = ["us-east-1a", "us-east-1b", "us-east-1c"]
}

# Sensitive variable
variable "api_token" {
  description = "API authentication token"
  type        = string
  default     = "example-token-12345"
  sensitive   = true
}

# Nullable variable
variable "custom_domain" {
  description = "Custom domain name (optional)"
  type        = string
  default     = null
  nullable    = true
}

