# Outputs demonstrating different features and use cases

# Simple string output
output "deployment_id" {
  description = "Unique deployment identifier"
  value       = random_uuid.project_id.result
}

# Sensitive output
output "internal_token" {
  description = "Internal authentication token"
  value       = random_string.suffix.result
  sensitive   = true
}

# Output with depends_on
output "config_ready" {
  description = "Indicates configuration is ready"
  value       = "Configuration file created successfully"
  depends_on  = [local_file.config]
}

# Complex output - object
output "deployment_info" {
  description = "Complete deployment information"
  value = {
    id          = random_uuid.project_id.result
    environment = var.environment
    project     = var.project_name
    file_count  = var.file_count
    suffix      = random_string.suffix.result
  }
}

# Output from data source
output "config_size_bytes" {
  description = "Size of configuration file in bytes"
  value       = length(data.local_file.config_read.content_base64)
}

# Conditional output
output "environment_message" {
  description = "Environment-specific message"
  value       = var.environment == "prod" ? "Production environment - handle with care!" : "Non-production environment"
}

# Output using for expression
output "file_details" {
  description = "Details of all created files"
  value = {
    for idx, file in local_file.demo_files :
    "file_${idx}" => {
      path = file.filename
      id   = file.id
    }
  }
}

# Output with precondition
output "validated_environment" {
  description = "Validated environment name"
  value       = var.environment
  
  precondition {
    condition     = contains(["dev", "staging", "prod"], var.environment)
    error_message = "Environment must be one of: dev, staging, prod"
  }
}

# JSON encoded output
output "config_json" {
  description = "Configuration as JSON string"
  value = jsonencode({
    project     = var.project_name
    environment = var.environment
    id          = random_uuid.project_id.result
  })
}

# List output
output "all_file_paths" {
  description = "All generated file paths"
  value = concat(
    [local_file.config.filename],
    [for f in local_file.demo_files : f.filename],
    [for k, f in local_file.environment_files : f.filename]
  )
}

