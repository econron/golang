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