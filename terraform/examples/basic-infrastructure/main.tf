# Basic Infrastructure Example
# Demonstrates core Terraform features

terraform {
  required_version = ">= 1.0"
  
  required_providers {
    local = {
      source  = "hashicorp/local"
      version = "~> 2.4"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.5"
    }
  }
}

# Provider configuration
provider "local" {}
provider "random" {}

# Variables - Input values for configuration
variable "environment" {
  description = "Environment name (dev, staging, prod)"
  type        = string
  default     = "dev"
}

variable "project_name" {
  description = "Name of the project"
  type        = string
  default     = "terraform-demo"
}

variable "file_count" {
  description = "Number of files to create"
  type        = number
  default     = 3
}

# Local values - Computed values used in configuration
locals {
  timestamp = timestamp()
  common_tags = {
    Environment = var.environment
    Project     = var.project_name
    ManagedBy   = "Terraform"
  }
  file_prefix = "${var.project_name}-${var.environment}"
}

# Resource: Random string generator
resource "random_string" "suffix" {
  length  = 8
  special = false
  upper   = false
}

# Resource: Generate random UUID
resource "random_uuid" "project_id" {}

# Resource: Create a configuration file
resource "local_file" "config" {
  filename = "${path.module}/generated/config.json"
  content = jsonencode({
    project_id  = random_uuid.project_id.result
    environment = var.environment
    created_at  = local.timestamp
    tags        = local.common_tags
  })
  
  file_permission = "0644"
}

# Resource: Create multiple files using count
resource "local_file" "demo_files" {
  count = var.file_count
  
  filename = "${path.module}/generated/${local.file_prefix}-file-${count.index}.txt"
  content  = <<-EOT
    File Number: ${count.index}
    Project: ${var.project_name}
    Environment: ${var.environment}
    Unique ID: ${random_string.suffix.result}
    Created: ${local.timestamp}
  EOT
  
  file_permission = "0644"
}

# Resource: Create files using for_each
resource "local_file" "environment_files" {
  for_each = toset(["dev", "staging", "prod"])
  
  filename = "${path.module}/generated/config-${each.key}.txt"
  content  = "Environment: ${each.key}\nProject ID: ${random_uuid.project_id.result}"
  
  file_permission = "0644"
}

# Data source: Read an existing file (after it's created)
data "local_file" "config_read" {
  filename = local_file.config.filename
  
  depends_on = [local_file.config]
}

# Outputs - Values exported from the configuration
output "project_id" {
  description = "Unique project identifier"
  value       = random_uuid.project_id.result
}

output "random_suffix" {
  description = "Random suffix for resources"
  value       = random_string.suffix.result
}

output "config_file_path" {
  description = "Path to the generated configuration file"
  value       = local_file.config.filename
}

output "created_files" {
  description = "List of created file paths"
  value       = [for f in local_file.demo_files : f.filename]
}

output "environment_configs" {
  description = "Map of environment configuration files"
  value       = { for k, f in local_file.environment_files : k => f.filename }
}

output "config_content_length" {
  description = "Length of configuration file content"
  value       = length(data.local_file.config_read.content)
}

output "all_tags" {
  description = "Common tags applied to resources"
  value       = local.common_tags
}

