# Provisioners Examples
# Demonstrates different types of provisioners

terraform {
  required_providers {
    local = {
      source  = "hashicorp/local"
      version = "~> 2.4"
    }
  }
}

resource "local_file" "provisioner_demo" {
  filename = "${path.module}/provisioner-test.txt"
  content  = "Testing provisioners"
  
  # Local-exec provisioner - runs command on machine running Terraform
  provisioner "local-exec" {
    command = "echo 'File created: ${self.filename}'"
  }
  
  # Local-exec with different interpreter
  provisioner "local-exec" {
    command     = "echo File ID: ${self.id}"
    interpreter = ["/bin/bash", "-c"]
  }
  
  # Local-exec with working directory
  provisioner "local-exec" {
    command     = "ls -la"
    working_dir = path.module
  }
  
  # Local-exec with environment variables
  provisioner "local-exec" {
    command = "echo $FILE_NAME created in $ENVIRONMENT"
    
    environment = {
      FILE_NAME   = self.filename
      ENVIRONMENT = "demo"
    }
  }
  
  # Destroy-time provisioner
  provisioner "local-exec" {
    when    = destroy
    command = "echo 'File is being destroyed'"
  }
  
  # Provisioner with failure handling
  provisioner "local-exec" {
    command     = "echo 'This might fail' && exit 0"
    on_failure  = continue  # Options: fail (default), continue
  }
}

# File provisioner example (commented as it requires remote connection)
# resource "null_resource" "file_provisioner_example" {
#   # File provisioner - copies files to remote resource
#   provisioner "file" {
#     source      = "local-file.txt"
#     destination = "/tmp/remote-file.txt"
#     
#     connection {
#       type     = "ssh"
#       user     = "admin"
#       password = var.password
#       host     = var.host
#     }
#   }
# }

# Remote-exec provisioner example (commented as it requires remote connection)
# resource "null_resource" "remote_exec_example" {
#   # Remote-exec provisioner - runs commands on remote resource
#   provisioner "remote-exec" {
#     inline = [
#       "sudo apt-get update",
#       "sudo apt-get install -y nginx",
#       "sudo systemctl start nginx"
#     ]
#     
#     connection {
#       type        = "ssh"
#       user        = "ubuntu"
#       private_key = file("~/.ssh/id_rsa")
#       host        = var.host
#     }
#   }
# }

