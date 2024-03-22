terraform {
  required_version = "1.7.5"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.40.0"
    }
  }
}

provider "aws" {
  region     = "eu-central-1"
  access_key = "AKIA3KK3GJL6YKC3QGHE"
  secret_key = "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"
}

data "aws_availability_zones" "available" {
  state = "available"
}

resource "aws_db_instance" "tutorial_database" {
  allocated_storage      = 10
  engine                 = "mysql"
  engine_version         = "8.0.36"
  instance_class         = "db.t3.micro"
  db_name                = "tutorial"
  username               = var.db_username
  password               = var.db_password
  db_subnet_group_name   = aws_db_subnet_group.tutorial_db_subnet_group.id
  vpc_security_group_ids = [aws_security_group.tutorial_db_sg.id]
  skip_final_snapshot    = true
}

resource "aws_instance" "tutorial_web" {
  count                  = var.settings.web_app.count
  ami                    = "ami-0183b16fc359a89dd"
  instance_type          = "t2.micro"
  subnet_id              = aws_subnet.tutorial_public_subnet[count.index].id
  key_name               = aws_key_pair.tutorial_kp.key_name
  vpc_security_group_ids = [aws_security_group.tutorial_web_sg.id]
}
