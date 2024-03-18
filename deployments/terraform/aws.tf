terraform {
  cloud {
    organization = "golden-gecko"

    workspaces {
      name = "go-demo"
    }
  }

  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "5.40.0"
    }

    cloudflare = {
      source = "cloudflare/cloudflare"
      version = "4.26.0"
    }

    google = {
      source  = "hashicorp/google"
      version = "5.20.0"
    }

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.26.0"
    }
  }
}

provider "aws" {
  region = "eu-central-1"
  access_key = "AKIA3KK3GJL6YKC3QGHE"
  secret_key = "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"
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "2.77.0"

  name                 = "education"
  cidr                 = "10.0.0.0/16"
  azs                  = data.aws_availability_zones.available.names
  public_subnets       = ["10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24"]
  enable_dns_hostnames = true
  enable_dns_support   = true
}

resource "aws_instance" "app_server" {
  ami           = "ami-01be94ae58414ab2e"
  instance_type = "t2.micro"
}

resource "aws_db_instance" "app_db" {
  allocated_storage     = 10
  engine                = "mysql"
  instance_class        = "db.t3.micro"
  username              = "go_user"
  password              = "go_password"
  skip_final_snapshot   = true
}
