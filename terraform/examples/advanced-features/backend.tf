# Backend Configuration Examples
# Demonstrates different backend types for state storage

# Local backend (default)
# State is stored in a local file terraform.tfstate
terraform {
  backend "local" {
    path = "terraform.tfstate"
  }
}

# Uncomment to use S3 backend (AWS)
# terraform {
#   backend "s3" {
#     bucket         = "my-terraform-state"
#     key            = "terraform.tfstate"
#     region         = "us-east-1"
#     encrypt        = true
#     dynamodb_table = "terraform-state-lock"
#   }
# }

# Uncomment to use Azure Storage backend
# terraform {
#   backend "azurerm" {
#     resource_group_name  = "terraform-state-rg"
#     storage_account_name = "terraformstate"
#     container_name       = "tfstate"
#     key                  = "terraform.tfstate"
#   }
# }

# Uncomment to use Google Cloud Storage backend
# terraform {
#   backend "gcs" {
#     bucket = "my-terraform-state"
#     prefix = "terraform/state"
#   }
# }

# Uncomment to use Terraform Cloud/Enterprise
# terraform {
#   backend "remote" {
#     organization = "my-org"
#     
#     workspaces {
#       name = "my-workspace"
#     }
#   }
# }

