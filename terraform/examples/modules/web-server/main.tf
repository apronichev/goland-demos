# Example Terraform Module: Web Server
# Demonstrates module creation and reusability

variable "server_name" {
  description = "Name of the web server"
  type        = string
}

variable "port" {
  description = "Port number for the web server"
  type        = number
  default     = 8080
}

variable "environment" {
  description = "Environment name"
  type        = string
}

variable "enable_ssl" {
  description = "Enable SSL/TLS"
  type        = bool
  default     = false
}

locals {
  server_config = {
    name        = var.server_name
    port        = var.port
    environment = var.environment
    ssl_enabled = var.enable_ssl
    created_at  = timestamp()
  }
}

resource "random_uuid" "server_id" {}

resource "local_file" "server_config" {
  filename = "${path.module}/server-${var.server_name}.json"
  content  = jsonencode(local.server_config)
  
  file_permission = "0644"
}

output "server_id" {
  description = "Unique server identifier"
  value       = random_uuid.server_id.result
}

output "server_config_path" {
  description = "Path to server configuration file"
  value       = local_file.server_config.filename
}

output "server_endpoint" {
  description = "Server endpoint"
  value       = "${var.enable_ssl ? "https" : "http"}://${var.server_name}:${var.port}"
}

