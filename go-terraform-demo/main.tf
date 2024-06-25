//The AWS provider and the region where resources will be deployed. The region is set using a variable.
provider "aws" {
  region = var.aws_region
}

# Variables. These variables allow you to parameterize your configuration, making it more flexible and reusable.
variable "aws_region" {
  description = "The AWS region to deploy resources in"
  default     = "us-west-2"
}

variable "bucket_name" {
  description = "The name of the S3 bucket"
  default     = "my-terraform-bucket-example"
}

variable "instance_type" {
  description = "The EC2 instance type"
  default     = "t2.micro"
}

variable "key_name" {
  description = "The name of the EC2 key pair"
  default     = "my-key-pair"
}

# S3 Bucket. This creates an S3 bucket with a name specified by the bucket_name variable and tags it appropriately.
resource "aws_s3_bucket" "example" {
  bucket = var.bucket_name
  acl    = "private"

  tags = {
    Name        = "MyTerraformBucket"
    Environment = "Dev"
  }
}

# EC2 Instance. This creates an EC2 instance using an Amazon Linux 2 AMI, with the instance type and key name specified by variables.
resource "aws_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0"  # Amazon Linux 2 AMI ID
  instance_type = var.instance_type
  key_name      = var.key_name

  tags = {
    Name = "MyTerraformInstance"
  }
}

# Outputs
output "bucket_id" {
  description = "The ID of the S3 bucket"
  value       = aws_s3_bucket.example.id
}

output "instance_id" {
  description = "The ID of the EC2 instance"
  value       = aws_instance.example.id
}
