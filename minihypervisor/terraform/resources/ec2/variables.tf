# variables.tf
variable "name" {
  description = "module name"
  type        = string
  default     = "ec2 testing module"
}

variable "security_group_id" {
  description = "The security group ID to associate with the EC2 instance"
  type        = string
}

variable "subnet_id" {
  description = "subnet id"
  type = string
  sensitive = true
}

variable "key_name" {
  description = "key_name"
  type = string
  sensitive = true
}

