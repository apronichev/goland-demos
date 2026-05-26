# Lifecycle Rules Examples
# Demonstrates lifecycle meta-arguments

terraform {
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

provider "local" {}
provider "random" {}

# create_before_destroy - Creates new resource before destroying old one
resource "random_string" "cbd_example" {
  length  = 16
  special = false
  
  lifecycle {
    create_before_destroy = true
  }
}

resource "local_file" "cbd_file" {
  filename = "${path.module}/cbd-file.txt"
  content  = "String: ${random_string.cbd_example.result}"
  
  lifecycle {
    create_before_destroy = true
  }
}

# prevent_destroy - Prevents resource from being destroyed
# resource "local_file" "protected_file" {
#   filename = "${path.module}/protected.txt"
#   content  = "This file is protected from destruction"
#   
#   lifecycle {
#     prevent_destroy = true
#   }
# }

# ignore_changes - Ignores changes to specific attributes
resource "local_file" "ignore_changes_example" {
  filename = "${path.module}/ignore-changes.txt"
  content  = "Initial content"
  
  lifecycle {
    ignore_changes = [
      content,  # Changes to content will be ignored
    ]
  }
}

# Ignore all changes
resource "local_file" "ignore_all_example" {
  filename = "${path.module}/ignore-all.txt"
  content  = "Content that won't be updated by Terraform"
  
  lifecycle {
    ignore_changes = all
  }
}

# replace_triggered_by - Forces replacement when specified resource changes
resource "random_uuid" "version" {}

resource "local_file" "versioned_file" {
  filename = "${path.module}/versioned-file.txt"
  content  = "Version: ${random_uuid.version.result}"
  
  lifecycle {
    replace_triggered_by = [
      random_uuid.version.result
    ]
  }
}

# precondition and postcondition
variable "file_size_limit" {
  description = "Maximum file size in bytes"
  type        = number
  default     = 1000
}

resource "local_file" "validated_file" {
  filename = "${path.module}/validated.txt"
  content  = "Validated content"
  
  lifecycle {
    # Precondition - checked before resource operations
    precondition {
      condition     = var.file_size_limit > 0
      error_message = "File size limit must be greater than 0"
    }
    
    # Postcondition - checked after resource operations
    postcondition {
      condition     = self.id != ""
      error_message = "File must have a valid ID after creation"
    }
  }
}

# Combined lifecycle rules
resource "random_string" "combined_example" {
  length  = 12
  special = true
  
  keepers = {
    timestamp = timestamp()
  }
  
  lifecycle {
    create_before_destroy = true
    ignore_changes = [
      keepers,
    ]
  }
}

