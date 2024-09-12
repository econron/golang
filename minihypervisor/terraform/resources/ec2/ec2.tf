resource "aws_instance" "kvmec2" {
    ami = "ami-0cab37bd176bb80d3" // ubuntu
    instance_type = "z1d.metal"

    vpc_security_group_ids = [var.security_group_id]

    subnet_id = var.subnet_id // northeast-1a public

    associate_public_ip_address = true

    key_name = var.key_name

    tags = {
        Name = "kvmgo"
    }
}