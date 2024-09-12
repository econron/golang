module "testing-sg" {
    source = "./resources/sg"
    name = "testing module"
    vpc_id = var.vpc_id
    my_ip = var.my_ip
}

module "testing-ec2" {
    source = "./resources/ec2"
    security_group_id = module.testing-sg.security_group_id
    subnet_id = var.subnet_id
    key_name = var.key_name
}