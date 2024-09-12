# variables.tf
variable "name" {
  description = "module name"
  type        = string
  default     = "sg testing module"
}

variable "vpc_id" {
  description = "vpc id"
  type = string
  sensitive = true
}

variable "my_ip" {
  description = "my ip"
  type = string
  sensitive = true
}