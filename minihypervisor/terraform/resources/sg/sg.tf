resource "aws_security_group" "for_my_testing_ec2" {
    name = "testing_sg"
    description = "this is a test sg"
    vpc_id = var.vpc_id

    ingress {
        description = "SSH access"
        from_port = 22
        to_port = 22
        protocol = "tcp"
        cidr_blocks = [var.my_ip]
    }

    ingress {
        description = "http public but restricted access"
        from_port = 80
        to_port = 80
        protocol = "tcp"
        cidr_blocks = [var.my_ip]
    }

    ingress {
        description = "https public but restricted access"
        from_port = 443
        to_port = 443
        protocol = "tcp"
        cidr_blocks = [var.my_ip]
    }

    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }

    tags = {
        Name = "kvmgo"
    }
}