# Example of using modules
# Demonstrates how to call and reuse modules

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

provider "local" {}
provider "random" {}

# Using the web-server module for development
module "dev_web_server" {
  source = "./web-server"
  
  server_name = "dev-app-server"
  port        = 8080
  environment = "development"
  enable_ssl  = false
}

# Using the web-server module for production
module "prod_web_server" {
  source = "./web-server"
  
  server_name = "prod-app-server"
  port        = 443
  environment = "production"
  enable_ssl  = true
}

# Using the same module with different configuration
module "api_server" {
  source = "./web-server"
  
  server_name = "api-server"
  port        = 3000
  environment = "production"
  enable_ssl  = true
}

# Outputs from modules
output "dev_server_id" {
  description = "Development server ID"
  value       = module.dev_web_server.server_id
}

output "dev_server_endpoint" {
  description = "Development server endpoint"
  value       = module.dev_web_server.server_endpoint
}

output "prod_server_id" {
  description = "Production server ID"
  value       = module.prod_web_server.server_id
}

output "prod_server_endpoint" {
  description = "Production server endpoint"
  value       = module.prod_web_server.server_endpoint
}

output "api_server_endpoint" {
  description = "API server endpoint"
  value       = module.api_server.server_endpoint
}

output "all_server_ids" {
  description = "All server IDs"
  value = {
    dev  = module.dev_web_server.server_id
    prod = module.prod_web_server.server_id
    api  = module.api_server.server_id
  }
}

