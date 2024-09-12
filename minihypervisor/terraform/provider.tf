terraform {
    required_providers {
      aws = {
        source = "hashicorp/aws"
        version = "~> 4.16"
      }
    }
    backend "s3" {
        bucket = "tfstate-manager-test"
        region = "ap-northeast-1"
        key = "kvmapi"
        encrypt = true
    }

    required_version = ">= 1.2.0"
}

provider "aws" {
  region = "ap-northeast-1"
}